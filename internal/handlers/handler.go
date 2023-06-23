package handlers

import (
	"bank-api/internal/core/ports"
	"bank-api/pkg/logger"
)

type BankHandler struct {
	service ports.BankServicePort
	log     *logger.Logger
}

func InitHandler(service ports.BankServicePort, log *logger.Logger) *BankHandler {
	return &BankHandler{
		service: service,
		log:     log,
	}
}
