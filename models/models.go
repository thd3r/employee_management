package models

import (
	"github.com/thd3r/employee_management/internal/database"
	"github.com/thd3r/employee_management/models/employee"

	"gorm.io/gorm"
)

func Health() map[string]string {
	return database.DbInstance.Health()
}

func CreateEmployee(employee *employee.Employee) *gorm.DB {
	return database.DbInstance.Cursor.Create(&employee)
}

func UpdateEmployee(employee *employee.Employee, employeeId string) *gorm.DB {
	return database.DbInstance.Cursor.Where("id = ?", employeeId).Updates(&employee)
}

func DeleteEmployee(employee *employee.Employee, employeeId string) *gorm.DB {
	return database.DbInstance.Cursor.Where("id = ?", employeeId).Delete(&employee)
}

func GetAllEmployee(employee *[]employee.Employee) *gorm.DB {
	return database.DbInstance.Cursor.Find(&employee)
}

func GetEmployeeById(employee *employee.Employee, employeeId string) *gorm.DB {
	return database.DbInstance.Cursor.First(&employee, "id = ?", employeeId)
}