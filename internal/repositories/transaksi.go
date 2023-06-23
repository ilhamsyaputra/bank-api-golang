package repositories

import (
	"bank-api/internal/data"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func (b *BankRepository) AddMutasiTransaksi(tx *sqlx.Tx, data data.Transaksi) (err error) {
	b.log.Info(
		logrus.Fields{}, nil, "Execute: BankRepository.AddMutasiTransaksi started",
	)

	params := map[string]interface{}{
		"id":             data.Id,
		"no_rekening":    data.NoRekening,
		"kode_transaksi": data.KodeTransaksi,
		"nominal":        data.Nominal,
	}

	query := "insert into transaksi (id, no_rekening, kode_transaksi, nominal) values (:id, :no_rekening, :kode_transaksi, :nominal)"

	// START insert to db
	_, err = tx.NamedExec(query, params)
	if err != nil {
		return
	}
	// -- END insert to DB

	return
}
