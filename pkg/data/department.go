package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Department struct {
	Code int    `gorm:"code"`
	Name string `gorm:"name"`
}

type DepartmentData struct {
	db *gorm.DB
}

func NewDepartmentData(db *gorm.DB) *DepartmentData {
	return &DepartmentData{db: db}
}

func (u DepartmentData) Add(department Department) (int, error) {
	result := u.db.Create(&department)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create department, error: %w", result.Error)
	}
	return department.Code, nil
}

func (u DepartmentData) Read(code int) (Department, error) {
	var department Department
	result := u.db.Find(&department, code)
	if result.Error != nil {
		return department, fmt.Errorf("can't read department with given id, error: %w", result.Error)
	}
	return department, nil
}

func (u DepartmentData) ReadAll() ([]Department, error) {
	var departments []Department
	result := u.db.Find(&departments)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read departments from database, error: %w", result.Error)
	}
	return departments, nil
}

func (u DepartmentData) ChangeName(code int, name string) (int, error) {
	result := u.db.Model(Department{}).Where("code = ", code).Update("name", name)
	if result.Error != nil {
		return -1, fmt.Errorf("can't update name, error: %w", result.Error)
	}
	return code, nil
}

func (u DepartmentData) Delete(code int) error {
	result := u.db.Delete(&Department{}, code)
	if result.Error != nil {
		return fmt.Errorf("can't delete department from database, error: %w", result.Error)
	}
	return nil
}
