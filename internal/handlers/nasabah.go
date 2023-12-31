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
				"status":  "error",
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
				"status":  "error",
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

func (h *BankHandler) Tabung(c *fiber.Ctx) error {
	var requestPayload data.TrxRequest

	err := c.BodyParser(&requestPayload)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ERROR OCCURED")

		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	saldo, err := h.service.Tabung(requestPayload)
	if err != nil && err.Error() == "INVALID" {
		return c.Status(400).JSON(
			fiber.Map{
				"status": "error",
				"remark": "Tidak dapat melakukan transaksi tabung. Nomor rekening tidak valid.",
			},
		)
	}
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"saldo": saldo,
			},
		},
	)
}

func (h *BankHandler) Tarik(c *fiber.Ctx) error {
	var requestPayload data.TrxRequest

	err := c.BodyParser(&requestPayload)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("ERROR OCCURED")

		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	saldo, err := h.service.Tarik(requestPayload)
	if err != nil && err.Error() == "INVALID" {
		return c.Status(400).JSON(
			fiber.Map{
				"status": "error",
				"remark": "Tidak dapat melakukan transaksi tarik. Nomor rekening tidak valid.",
			},
		)
	}
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"saldo": saldo,
			},
		},
	)
}

func (h *BankHandler) GetSaldo(c *fiber.Ctx) error {
	no_rekening := c.Params("no_rekening")

	saldo, err := h.service.GetSaldo(no_rekening)
	if err != nil && err.Error() == "INVALID" {
		return c.Status(400).JSON(
			fiber.Map{
				"status": "error",
				"remark": "Tidak mendapatkan saldo. Nomor rekening tidak valid.",
			},
		)
	}
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"status": "success",
			"data": fiber.Map{
				"saldo": saldo,
			},
		},
	)
}

func (h *BankHandler) GetMutasi(c *fiber.Ctx) error {
	no_rekening := c.Params("no_rekening")

	mutasi, err := h.service.GetMutasi(no_rekening)
	if err != nil && err.Error() == "INVALID" {
		return c.Status(400).JSON(
			fiber.Map{
				"status": "error",
				"remark": "Tidak mendapatkan data mutasi. Nomor rekening tidak valid.",
			},
		)
	}
	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"status":  "error",
				"message": err,
			},
		)
	}

	return c.Status(200).JSON(
		fiber.Map{
			"status": "success",
			"data":   mutasi,
		},
	)
}
