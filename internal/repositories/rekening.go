package repositories

import (
	"bank-api/pkg/utils"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (b *BankRepository) RegisterRekening(tx *sqlx.Tx, no_nasabah int) (no_rekening string, err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: BankRepository.RegisterRekening started",
	)

	no_rekening = utils.GenerateNoRekening()

	params := map[string]interface{}{
		"no_nasabah":  no_nasabah,
		"no_rekening": no_rekening,
	}

	query := "insert into rekening values (:no_rekening, :no_nasabah, 0) RETURNING no_rekening"

	// START insert to db
	_, err = tx.NamedExec(query, params)
	if err != nil {
		return "", err
	}

	// err = tx.Commit()
	// if err != nil {
	// 	b.log.Error(
	// 		logrus.Fields{"error": err.Error()}, nil, "failed to commit transaction",
	// 	)
	// }
	// -- END insert to DB
	fmt.Println(no_rekening)

	return
}