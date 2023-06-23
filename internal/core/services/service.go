package services

import (
	"bank-api/internal/core/ports"
	"bank-api/pkg/logger"
)

type BankService struct {
	repository ports.BankRepositoryPort
	log        *logger.Logger
}

func InitBankService(repository ports.BankRepositoryPort, log *logger.Logger) *BankService {
	return &BankService{
		repository: repository,
		log:        log,
	}
}
