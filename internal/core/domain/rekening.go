package domain

type Rekening struct {
	NoRekening int `db:"no_rekening" json:"no_rekening"`
	NoNasabah  int `db:"no_nasabah" json:"no_nasabah"`
	Saldo      int `db:"saldo" json:"saldo"`
}

func NewRekening(noRekening int, noNasabah int, saldo int) *Rekening {
	return &Rekening{
		NoRekening: noRekening,
		NoNasabah:  noNasabah,
		Saldo:      saldo,
	}
}
