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

func (l LecturerData) Add(lecturer Lecturer) (int, error) {
	result := l.db.Create(&lecturer)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create lecturer, error: %w", result.Error)
	}
	return lecturer.Id, nil
}

func (l LecturerData) Read(id int) (Lecturer, error) {
	var lecturer Lecturer
	result := l.db.Find(&lecturer, id)
	if result.Error != nil {
		return lecturer, fmt.Errorf("can't read lecturer with given id, error: %w", result.Error)
	}
	return lecturer, nil
}

func (l LecturerData) ReadAll() ([]Lecturer, error) {
	var lecturers []Lecturer
	result := l.db.Find(&lecturers)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read lecturers from database, error: %w", result.Error)
	}
	return lecturers, nil
}

func (l LecturerData) ChangeFullName(id int, firstName, lastName string) (int, error) {
	result := l.db.Model(Lecturer{}).Where("id = ?", id).Updates(Lecturer{FirstName: firstName, LastName: lastName})
	if result.Error != nil {
		return -1, fmt.Errorf("can't update name, error: %w", result.Error)
	}
	return id, nil
}

func (l LecturerData) Delete(id int) error {
	result := l.db.Delete(&Lecturer{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't delete lecturer from database, error: %w", result.Error)
	}
	return nil
}
