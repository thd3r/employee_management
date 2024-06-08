package controllers

import (
	"fmt"
	"time"
	"strings"

	"github.com/thd3r/employee_management/models"
	"github.com/thd3r/employee_management/models/employe"
	"github.com/thd3r/employee_management/models/schema"
	"github.com/thd3r/employee_management/utils"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code": 200,
	})
}

func Health(c *fiber.Ctx) error {
	dataHealth := models.Health()
	return c.JSON(dataHealth)
}

func CreateEmployeHandler(c *fiber.Ctx) error {
	var dataEmployeCreate schema.CreateEmployeSchema

	if err := c.BodyParser(&dataEmployeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest,
		})
	}

	if err := utils.ValidateRequest(dataEmployeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	newEmploye := employe.Employe{
		Id:        uuid.New().String(),
		Name:      dataEmployeCreate.Name,
		Role:      dataEmployeCreate.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data := models.CreateEmploye(&newEmploye)
	if data.Error != nil {
		if strings.Contains(data.Error.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"msg":  fmt.Sprintf("Employee named %s is available", dataEmployeCreate.Name),
				"code": fiber.StatusConflict,
			})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"msg":  data.Error.Error(),
				"code": fiber.StatusBadGateway,
			})
		}
	} else if data.RowsAffected < 1 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"msg":  data.Error.Error(),
			"code": fiber.StatusBadGateway,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newEmploye,
		"code": fiber.StatusCreated,
	})
}

func GetAllEmployeHandler(c *fiber.Ctx) error {
	var employe []employe.Employe

	models.GetAllEmploye(&employe)
	return c.Status(fiber.StatusOK).JSON(employe)
}

func GetEmployeByIdHandler(c *fiber.Ctx) error {
	employeId := c.Params("employeId")

	var dataEmploye employe.Employe
	if result := models.GetEmployeById(&dataEmploye, employeId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg":  fmt.Sprintf("Employee with id %v not found", employeId),
				"code": fiber.StatusNotFound,
			})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"msg":  result.Error.Error(),
			"code": fiber.StatusBadGateway,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataEmploye,
		"code": fiber.StatusOK,
	})
}

func UpdateEmployeHandler(c *fiber.Ctx) error {
	employeId := c.Params("employeId")
	var dataEmployeUpdate schema.UpdateEmployeSchema

	if err := c.BodyParser(&dataEmployeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest,
		})
	}

	if err := utils.ValidateRequest(dataEmployeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	var dataEmploye employe.Employe

	if result := models.GetEmployeById(&dataEmploye, employeId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"msg":  fmt.Sprintf("Employee with name %s not found", dataEmployeUpdate.Name),
				"code": fiber.StatusNotFound,
			})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"msg":  result.Error.Error(),
			"code": fiber.StatusBadGateway,
		})
	}

	dataEmploye.Name = dataEmployeUpdate.Name
	dataEmploye.Role = dataEmployeUpdate.Role
	dataEmploye.UpdatedAt = time.Now()

	if result := models.UpdateEmploye(&dataEmploye, employeId); result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"msg":  fmt.Sprintf("Employee named %s is available", dataEmployeUpdate.Name),
				"code": fiber.StatusConflict,
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":  result.Error.Error(),
				"code": fiber.StatusBadRequest,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataEmploye,
		"code": fiber.StatusOK,
	})

}

func DeleteEmployeHandler(c *fiber.Ctx) error {
	employeId := c.Params("employeId")
	result := models.DeleteEmploye(&employe.Employe{}, employeId)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"msg":  result.Error.Error(),
			"code": fiber.StatusBadGateway,
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"msg":  fmt.Sprintf("Employee with id %v not found", employeId),
			"code": fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  fmt.Sprintf("Employee with id %v has been deleted", employeId),
		"code": fiber.StatusOK,
	})
}
