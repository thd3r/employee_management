package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thd3r/employee_management/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "Thd3rServer",
			AppName:      "Employee Management",
		}),

		db: database.New(),
	}

	return server
}
