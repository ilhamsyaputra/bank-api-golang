package data

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
