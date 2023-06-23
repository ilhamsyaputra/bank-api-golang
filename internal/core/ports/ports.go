package ports

import (
	"bank-api/internal/data"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type BankServicePort interface {
	Register(requestPayload data.RegisterRequest) (no_rekening string, err error)
}

type BankRepositoryPort interface {
	Begin() (tx *sqlx.Tx, err error)
	Rollback(tx *sqlx.Tx)
	Commit(tx *sqlx.Tx)
	IsNasabahExist(tx *sqlx.Tx, requestPayload data.RegisterRequest) (isExist bool, err error)
	RegisterNasabah(tx *sqlx.Tx, requestPayload data.RegisterRequest) (no_rekening int, err error)
	RegisterRekening(tx *sqlx.Tx, no_nasabah int) (no_rekening string, err error)
}

type BankHandlersPort interface {
	Register(c *fiber.Ctx) error
}
