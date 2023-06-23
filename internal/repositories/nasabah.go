package repositories

import (
	"bank-api/internal/data"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (b *BankRepository) IsNasabahExist(tx *sqlx.Tx, requestPayload data.RegisterRequest) (isExist bool, err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: IsNasabahExist started",
	)

	params := map[string]interface{}{
		"nik":   requestPayload.Nik,
		"no_hp": requestPayload.Nohp,
	}

	query := "SELECT COUNT(*) FROM nasabah WHERE nik = :nik OR no_hp = :no_hp"

	var count int
	prepareQuery, err := tx.PrepareNamed(query)
	if err != nil {
		return false, err
	}

	err = prepareQuery.Get(&count, params)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

func (b *BankRepository) RegisterNasabah(tx *sqlx.Tx, requestPayload data.RegisterRequest) (no_nasabah int, err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: BankRepository.Register started",
	)

	params := map[string]interface{}{
		"nama":  requestPayload.Nama,
		"nik":   requestPayload.Nik,
		"no_hp": requestPayload.Nohp,
	}

	query := "insert into nasabah (nama, nik, no_hp) values (:nama, :nik, :no_hp) RETURNING no_nasabah"

	// START insert to db
	_, err = tx.NamedExec(query, params)
	if err != nil {
		return 0, err
	}
	// -- END insert to DB

	err = tx.QueryRow("SELECT last_value FROM nasabah_no_nasabah_seq").Scan(&no_nasabah)
	if err != nil {
		return 0, err
	}

	return
}
