package domain

import "time"

type Nasabah struct {
	NoNasabah         string    `db:"no_rekening" json:"no_rekening"`
	Nama              string    `db:"nama" json:"nama"`
	Nik               string    `db:"nik" json:"nik"`
	NoHp              string    `db:"no_hp" json:"no_hp"`
	TanggalRegistrasi time.Time `db:"tanggal_registrasi" json:"tanggal_registrasi"`
}

func NewNasabah(noNasabah string, nama string, nik string, nohp string) *Nasabah {
	return &Nasabah{
		NoNasabah:         noNasabah,
		Nama:              nama,
		Nik:               nik,
		NoHp:              nohp,
		TanggalRegistrasi: time.Now(),
	}
}
