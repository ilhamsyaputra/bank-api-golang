package repositories

import (
	"bank-api/pkg/logger"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BankRepository struct {
	db  *sqlx.DB
	log *logger.Logger
}

func (b *BankRepository) Begin() (tx *sqlx.Tx, err error) {
	tx, err = b.db.Beginx()
	if err != nil {
		b.log.Error(
			logrus.Fields{"error": err.Error()}, nil, "failed to start transaction",
		)
	}
	return
}

func (b *BankRepository) Rollback(tx *sqlx.Tx) {
	err := tx.Rollback()
	if err != nil {
		b.log.Error(
			logrus.Fields{"error": err.Error()}, nil, "failed to rollback transaction",
		)
	}
}

func (b *BankRepository) Commit(tx *sqlx.Tx) {
	err := tx.Commit()
	if err != nil {
		b.log.Error(
			logrus.Fields{"error": err.Error()}, nil, "failed to commit transaction",
		)
	}
}

func InitRepository(driver, host, user, password, database string, port int, log *logger.Logger) *BankRepository {
	address := fmt.Sprintf(`user=%s password=%s dbname=%s`, user, password, database)

	db, err := sqlx.Connect(driver, address)
	if err != nil {
		panic(err)
	}
	return &BankRepository{
		db:  db,
		log: log,
	}
}
