package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Course struct {
	Id          int    `gorm:"id"`
	Title       string `gorm:"title"`
	Description string `gorm:"description"`
	Fee         int    `gorm:"fee"`
}

type CourseData struct {
	db *gorm.DB
}

func NewCourseData(db *gorm.DB) *CourseData {
	return &CourseData{db: db}
}

func (d CourseData) Add(course Course) (int, error) {
	result := d.db.Create(&course)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create course, error: %w", result.Error)
	}
	return course.Id, nil
}

func (d CourseData) Read(id int) (Course, error) {
	var course Course
	result := d.db.Find(&course, id)
	if result.Error != nil {
		return course, fmt.Errorf("can't read course with given id, error: %w", result.Error)
	}
	return course, nil
}

func (d CourseData) ReadAll() ([]Course, error) {
	var courses []Course
	result := d.db.Find(&courses)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read courses from database, error: %w", result.Error)
	}
	return courses, nil
}

func (d CourseData) ChangeDescription(id int, description string) (int, error) {
	result := d.db.Model(Course{}).Where("id = ?", id).Update("description", description)
	if result.Error != nil {
		return -1, fmt.Errorf("can't update course description, error: %w", result.Error)
	}
	return id, nil
}

func (d CourseData) Delete(id int) error {
	result := d.db.Delete(&Course{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't delete course from database, error: %w", result.Error)
	}
	return nil
}
