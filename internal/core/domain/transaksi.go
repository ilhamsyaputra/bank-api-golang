package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaksi struct {
	Id               uuid.UUID `db:"id" json:"id"`
	NoRekening       int       `db:"no_rekening" json:"no_rekening"`
	KodeTransaksi    string    `db:"kode_tranksaksi" json:"kode_transaksi"`
	Nominal          int       `db:"nominal" json:"nominal"`
	TanggalTransaksi time.Time `db:"tanggal_transaksi" json:"tanggal_transaksi"`
}

func NewTransaksi(noRekening int, kodeTransaksi string, nominal int) *Transaksi {
	return &Transaksi{
		Id:               uuid.New(),
		NoRekening:       noRekening,
		KodeTransaksi:    kodeTransaksi,
		Nominal:          nominal,
		TanggalTransaksi: time.Now(),
	}
}
