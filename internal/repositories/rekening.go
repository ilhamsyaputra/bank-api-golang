package repositories

import (
	"bank-api/internal/data"
	"bank-api/pkg/utils"

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
	// -- END insert to DB

	return
}

func (b *BankRepository) IsRekeningValid(tx *sqlx.Tx, requestPayload data.TrxRequest) (valid bool, err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: IsRekeningValid started",
	)

	params := map[string]interface{}{
		"no_rekening": requestPayload.NoRekening,
	}

	query := "SELECT COUNT(*) FROM rekening WHERE no_rekening = :no_rekening"

	var count int
	prepareQuery, err := tx.PrepareNamed(query)
	if err != nil {
		return false, err
	}

	err = prepareQuery.Get(&count, params)
	if err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (b *BankRepository) GetSaldoByRekening(tx *sqlx.Tx, no_rekening string) (saldo int, err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: BankRepository.GetSaldoByRekening started",
	)

	query := "select saldo from rekening where no_rekening = $1"

	err = tx.Get(&saldo, query, no_rekening)
	return
}

func (b *BankRepository) AddSaldoByRekening(tx *sqlx.Tx, request data.TrxRequest) (err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: BankRepository.GetSaldoByRekening started",
	)

	params := map[string]interface{}{
		"no_rekening": request.NoRekening,
		"nominal":     request.Nominal,
	}

	query := "update rekening set saldo = :nominal where no_rekening = :no_rekening"

	_, err = tx.NamedExec(query, params)

	return
}

func (b *BankRepository) SubstractSaldoByRekening(tx *sqlx.Tx, request data.TrxRequest) (err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: BankRepository.SubstractSaldoByRekening started",
	)

	params := map[string]interface{}{
		"no_rekening": request.NoRekening,
		"nominal":     request.Nominal,
	}

	query := "update rekening set saldo = :nominal where no_rekening = :no_rekening"

	_, err = tx.NamedExec(query, params)

	return
}
