package handlers

import (
	"bank-api/internal/data"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

func (h *BankHandler) Register(c *fiber.Ctx) error {
	var requestPayload data.RegisterRequest

	err := c.BodyParser(&requestPayload)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ERROR OCCURED")

		return c.Status(500).JSON(
			fiber.Map{
				"status": "error",
				"message": err,
			},
		)
	}
	
	no_rekening, err := h.service.Register(requestPayload)
	if err != nil && err.Error() == "EXIST" {
		return c.Status(400).JSON(
			fiber.Map{
				"status": "error",
				"remark": "Tidak dapat registrasi nasabah baru. NIK atau Nomor HP sudah terdaftar",
			},
		)
	}
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status": "error",
				"message": err,
			},
		)
	}
	
	return c.Status(201).JSON(
		fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"no_rekening": no_rekening,
			},
		},
	)
}
