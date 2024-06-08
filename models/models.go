package models

import (
	"github.com/thd3r/employee_management/internal/database"
	"github.com/thd3r/employee_management/models/employe"

	"gorm.io/gorm"
)

func Health() map[string]string {
	return database.Health()
}

func CreateEmploye(employe *employe.Employe) *gorm.DB {
	return database.Instance.Cursor.Create(&employe)
}

func UpdateEmploye(employe *employe.Employe, employeId string) *gorm.DB {
	return database.Instance.Cursor.Where("id = ?", employeId).Updates(&employe)
}

func DeleteEmploye(employe *employe.Employe, employeId string) *gorm.DB {
	return database.Instance.Cursor.Where("id = ?", employeId).Delete(&employe)
}

func GetAllEmploye(employe *[]employe.Employe) *gorm.DB {
	return database.Instance.Cursor.Find(&employe)
}

func GetEmployeById(employe *employe.Employe, employeId string) *gorm.DB {
	return database.Instance.Cursor.First(&employe, "id = ?", employeId)
}