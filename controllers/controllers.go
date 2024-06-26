package controllers

import (
	"github.com/thd3r/employee_management/models"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": fiber.StatusOK,
	})
}

func IndexApiHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"errors":  `You don't have permission to access this resource`,
		"code": fiber.StatusForbidden,
	})
}

func Health(c *fiber.Ctx) error {
	dataHealth := models.Health()
	return c.Status(fiber.StatusOK).JSON(dataHealth)
}

func ApiSitemap(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(c.App().Stack())
}
