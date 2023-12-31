package ports

import (
	"bank-api/internal/data"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type BankServicePort interface {
	Register(requestPayload data.RegisterRequest) (no_rekening string, err error)
	Tabung(requestPayload data.TrxRequest) (saldo int, err error)
	Tarik(requestPayload data.TrxRequest) (saldo int, err error)
	GetSaldo(no_rekening string) (saldo int, err error)
	GetMutasi(no_rekening string) (response []data.Mutasi, err error)
}

type BankRepositoryPort interface {
	Begin() (tx *sqlx.Tx, err error)
	Rollback(tx *sqlx.Tx)
	Commit(tx *sqlx.Tx)
	IsNasabahExist(tx *sqlx.Tx, requestPayload data.RegisterRequest) (isExist bool, err error)
	RegisterNasabah(tx *sqlx.Tx, requestPayload data.RegisterRequest) (no_rekening int, err error)
	IsRekeningValid(tx *sqlx.Tx, no_rekening string) (valid bool, err error)
	RegisterRekening(tx *sqlx.Tx, no_nasabah int) (no_rekening string, err error)
	GetSaldoByRekening(tx *sqlx.Tx, no_rekening string) (saldo int, err error)
	AddSaldoByRekening(tx *sqlx.Tx, request data.TrxRequest) (err error)
	AddMutasiTransaksi(tx *sqlx.Tx, requestPayload data.Transaksi) (err error)
	SubstractSaldoByRekening(tx *sqlx.Tx, request data.TrxRequest) (err error)
	GetMutasiByRekening(tx *sqlx.Tx, no_rekening string) (response []data.Mutasi, err error)
}

type BankHandlersPort interface {
	Register(c *fiber.Ctx) error
	Tabung(c *fiber.Ctx) error
	Tarik(c *fiber.Ctx) error
	GetSaldo(c *fiber.Ctx) error
	GetMutasi(c *fiber.Ctx) error
}
