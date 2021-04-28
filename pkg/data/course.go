package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Course struct {
	Code           int    `gorm:"code"`
	Title          string `gorm:"title"`
	DepartmentCode int    `gorm:"department_code"`
	Description    string `gorm:"description"`
}

type CourseData struct {
	db *gorm.DB
}

func NewCourseData(db *gorm.DB) *CourseData {
	return &CourseData{db: db}
}

func (c CourseData) Add(course Course) (int, error) {
	result := c.db.Create(&course)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create course, error: %w", result.Error)
	}
	return course.Code, nil
}

func (c CourseData) Read(code int) (Course, error) {
	var course Course
	result := c.db.Find(&course, code)
	if result.Error != nil {
		return course, fmt.Errorf("can't read course with given id, error: %w", result.Error)
	}
	return course, nil
}

func (c CourseData) ReadAll() ([]Course, error) {
	var courses []Course
	result := c.db.Find(&courses)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read courses from database, error: %w", result.Error)
	}
	return courses, nil
}

func (c CourseData) ChangeDescription(code int, description string) (int, error) {
	result := c.db.Model(Course{}).Where("code = ?", code).Update("description", description)
	if result.Error != nil {
		return -1, fmt.Errorf("can't update course description, error: %w", result.Error)
	}
	return code, nil
}

func (c CourseData) Delete(code int) error {
	result := c.db.Delete(&Course{}, code)
	if result.Error != nil {
		return fmt.Errorf("can't delete course from database, error: %w", result.Error)
	}
	return nil
}

func (c CourseData) GetDepartmentName(code int) (string, error) {
	var departmentName string
	result := c.db.Model(&Course{}).
		Select("departments.name").
		Joins("join departments on department_code = departments.code").
		Where("courses.code = ?", code).
		Scan(&departmentName)
	if result.Error != nil {
		return "", fmt.Errorf("can't get department name from database, error: %w", result.Error)
	}
	return departmentName, nil
}
