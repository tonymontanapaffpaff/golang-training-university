package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Lecturer struct {
	Id             int    `gorm:"id"`
	FirstName      string `gorm:"first_name"`
	LastName       string `gorm:"last_name"`
	DepartmentCode int    `gorm:"department_code"`
}

type LecturerData struct {
	db *gorm.DB
}

func NewLecturerData(db *gorm.DB) *LecturerData {
	return &LecturerData{db: db}
}

func (u LecturerData) Add(lecturer Lecturer) (int, error) {
	result := u.db.Create(&lecturer)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create lecturer, error: %w", result.Error)
	}
	return lecturer.Id, nil
}

func (u LecturerData) Read(id int) (Lecturer, error) {
	var lecturer Lecturer
	result := u.db.Find(&lecturer, id)
	if result.Error != nil {
		return lecturer, fmt.Errorf("can't read lecturer with given id, error: %w", result.Error)
	}
	return lecturer, nil
}

func (u LecturerData) ReadAll() ([]Lecturer, error) {
	var lecturers []Lecturer
	result := u.db.Find(&lecturers)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read lecturers from database, error: %w", result.Error)
	}
	return lecturers, nil
}

func (u LecturerData) ChangeFullName(id int, firstName, lastName string) (int, error) {
	result := u.db.Model(Lecturer{}).Where("id = ?", id).Updates(Lecturer{FirstName: firstName, LastName: lastName})
	if result.Error != nil {
		return -1, fmt.Errorf("can't update name, error: %w", result.Error)
	}
	return id, nil
}

func (u LecturerData) Delete(id int) error {
	result := u.db.Delete(&Lecturer{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't delete lecturer from database, error: %w", result.Error)
	}
	return nil
}
