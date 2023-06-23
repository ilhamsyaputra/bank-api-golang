package ports

import (
	"bank-api/internal/data"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type BankServicePort interface {
	Register(requestPayload data.RegisterRequest) (no_rekening string, err error)
	Tabung(requestPayload data.TabungRequest) (saldo int, err error)
}

type BankRepositoryPort interface {
	Begin() (tx *sqlx.Tx, err error)
	Rollback(tx *sqlx.Tx)
	Commit(tx *sqlx.Tx)
	IsNasabahExist(tx *sqlx.Tx, requestPayload data.RegisterRequest) (isExist bool, err error)
	RegisterNasabah(tx *sqlx.Tx, requestPayload data.RegisterRequest) (no_rekening int, err error)
	IsRekeningValid(tx *sqlx.Tx, requestPayload data.TabungRequest) (valid bool, err error)
	RegisterRekening(tx *sqlx.Tx, no_nasabah int) (no_rekening string, err error)
	GetSaldoByRekening(tx *sqlx.Tx, no_rekening string) (saldo int, err error)
	AddSaldoByRekening(tx *sqlx.Tx, request data.TabungRequest) (err error)
	AddMutasiTransaksi(tx *sqlx.Tx, requestPayload data.Transaksi) (err error)
}

type BankHandlersPort interface {
	Register(c *fiber.Ctx) error
	Tabung(c *fiber.Ctx) error
}
