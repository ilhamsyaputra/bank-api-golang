package data

import "github.com/google/uuid"

type RegisterRequest struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	Nohp string `json:"no_hp"`
}

type RegisterResponse struct {
	NoRekening string `json:"no_rekening"`
}

type TabungRequest struct {
	NoRekening string `json:"no_rekening"`
	Nominal    int    `json:"nominal"`
}

type Transaksi struct {
	Id            uuid.UUID `json:"id_transaksi"`
	NoRekening    string    `json:"no_rekening"`
	KodeTransaksi string    `json:"kode_transaksi"`
	Nominal       int       `json:"nominal"`
}
