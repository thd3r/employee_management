package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/thd3r/employee_management/models"
	"github.com/thd3r/employee_management/models/employee"
	"github.com/thd3r/employee_management/models/employee/schema"
	"github.com/thd3r/employee_management/utils/validator"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateEmployeeHandler(c *fiber.Ctx) error {
	dataEmployeeCreate := new(schema.CreateEmployeeSchema)

	if err := c.BodyParser(&dataEmployeeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
			"code":   fiber.StatusBadRequest,
		})
	}

	if err := validator.ValidateRequestStruct(dataEmployeeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
			"code":   fiber.StatusBadRequest,
		})
	}

	newEmployee := employee.Employee{
		Id:        uuid.New().String(),
		Name:      dataEmployeeCreate.Name,
		Salary:    dataEmployeeCreate.Salary,
		Position:  dataEmployeeCreate.Position,
		CreatedAt: time.Now(),
	}

	data := models.CreateEmployee(&newEmployee)
	if data.Error != nil {
		if strings.Contains(data.Error.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"errors": fmt.Sprintf("Employee named %s is available", dataEmployeeCreate.Name),
				"code":   fiber.StatusConflict,
			})
		} else {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"errors": data.Error.Error(),
				"code":   fiber.StatusBadGateway,
			})
		}
	} else if data.RowsAffected < 1 {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"errors": data.Error.Error(),
			"code":   fiber.StatusBadGateway,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": newEmployee,
		"code": fiber.StatusCreated,
	})
}

func GetAllEmployeeHandler(c *fiber.Ctx) error {
	var dataEmployee []employee.Employee

	models.GetAllEmployee(&dataEmployee)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataEmployee,
		"code": fiber.StatusOK,
	})
}

func GetEmployeeByIdHandler(c *fiber.Ctx) error {
	employeId := c.Params("employeeId")
	var dataEmployee employee.Employee

	if result := models.GetEmployeeById(&dataEmployee, employeId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"errors": fmt.Sprintf("Employee with id %v not found", employeId),
				"code":   fiber.StatusNotFound,
			})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"errors": result.Error.Error(),
			"code":   fiber.StatusBadGateway,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataEmployee,
		"code": fiber.StatusOK,
	})
}

func UpdateEmployeeHandler(c *fiber.Ctx) error {
	employeId := c.Params("employeeId")
	dataEmployeeUpdate := new(schema.UpdateEmployeeSchema)

	if err := c.BodyParser(&dataEmployeeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
			"code":   fiber.StatusBadRequest,
		})
	}

	if err := validator.ValidateRequestStruct(dataEmployeeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err,
			"code":   fiber.StatusBadRequest,
		})
	}

	var dataEmployee employee.Employee

	if result := models.GetEmployeeById(&dataEmployee, employeId); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"errors": fmt.Sprintf("Employee with name %s not found", dataEmployeeUpdate.Name),
				"code":   fiber.StatusNotFound,
			})
		}

		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"errors": result.Error.Error(),
			"code":   fiber.StatusBadGateway,
		})
	}

	if dataEmployeeUpdate.Name != "" {
		dataEmployee.Name = dataEmployeeUpdate.Name
	}

	if dataEmployeeUpdate.Salary != "" {
		dataEmployee.Salary = dataEmployeeUpdate.Salary
	}

	if dataEmployeeUpdate.Position != "" {
		dataEmployee.Position = dataEmployeeUpdate.Position
	}

	dataEmployee.UpdatedAt = time.Now()

	if result := models.UpdateEmployee(&dataEmployee, employeId); result.Error != nil {
		if strings.Contains(result.Error.Error(), "Duplicate entry") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"errors": fmt.Sprintf("Employee named %s is available", dataEmployeeUpdate.Name),
				"code":   fiber.StatusConflict,
			})
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errors": result.Error.Error(),
				"code":   fiber.StatusBadRequest,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": dataEmployee,
		"code": fiber.StatusOK,
	})

}

func DeleteEmployeeHandler(c *fiber.Ctx) error {
	employeeId := c.Params("employeeId")
	result := models.DeleteEmployee(&employee.Employee{}, employeeId)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"errors": result.Error.Error(),
			"code":   fiber.StatusBadGateway,
		})
	}

	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"errors": fmt.Sprintf("Employee with id %v not found", employeeId),
			"code":   fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"errors": fmt.Sprintf("Employee with id %v has been deleted", employeeId),
		"code":   fiber.StatusOK,
	})
}
