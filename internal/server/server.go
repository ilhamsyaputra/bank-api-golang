package server

import (
	"bank-api/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	bankHandler ports.BankHandlersPort
}

func InitServer(handlers ports.BankHandlersPort) *Server {
	return &Server{
		bankHandler: handlers,
	}
}

func (s *Server) Start() {
	app := fiber.New()

	routes := app.Group("/v1")

	routes.Post("/daftar", s.bankHandler.Register)
	routes.Put("/tabung", s.bankHandler.Tabung)

	err := app.Listen(":2525")
	if err != nil {
		panic(err)
	}
}
