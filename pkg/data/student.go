package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	Id        int    `gorm:"id"`
	FirstName string `gorm:"first_name"`
	LastName  string `gorm:"last_name"`
	IsActive  bool   `gorm:"is_active"`
}

type StudentData struct {
	db *gorm.DB
}

func NewStudentData(db *gorm.DB) *StudentData {
	return &StudentData{db: db}
}

func (s StudentData) Add(student Student) (int, error) {
	result := s.db.Create(&student)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create student, error: %w", result.Error)
	}
	return student.Id, nil
}

func (s StudentData) Read(id int) (Student, error) {
	var student Student
	result := s.db.Find(&student, id)
	if result.Error != nil {
		return student, fmt.Errorf("can't read student with given id, error: %w", result.Error)
	}
	return student, nil
}

func (s StudentData) ReadAll() ([]Student, error) {
	var students []Student
	result := s.db.Find(&students)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read students from database, error: %w", result.Error)
	}
	return students, nil
}

func (s StudentData) ChangeStatus(id int) (int, error) {
	result := s.db.Exec("UPDATE students SET is_active = NOT is_active WHERE id = ?", id)
	if result.Error != nil {
		return -1, fmt.Errorf("can't update student status, error: %w", result.Error)
	}
	return id, nil
}

func (s StudentData) Delete(id int) error {
	result := s.db.Delete(&Student{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't delete student from database, error: %w", result.Error)
	}
	return nil
}

func (s StudentData) GetCurrentRate(id int) (float64, error) {
	var avg float64
	result := s.db.Model(&Student{}).
		Select("AVG(enrollments.average_grade)").
		Joins("join enrollments on id = enrollments.student_id").
		Where("students.id = ?", id).
		Scan(&avg)
	if result.Error != nil {
		return -1, fmt.Errorf("can't get list of courses from database, error: %w", result.Error)
	}
	return avg, nil
}

func (s StudentData) GetCoursesList(id int) ([]Course, error) {
	var courses []Course
	result := s.db.Model(&Student{}).
		Select("courses.code, courses.title, courses.department_code, courses.description").
		Joins("join enrollments on id = enrollments.student_id").
		Joins("join lessons on enrollments.lesson_id = lessons.id").
		Joins("join courses on lessons.course_code = courses.code").
		Where("students.id = ?", id).
		Scan(&courses)
	if result.Error != nil {
		return nil, fmt.Errorf("can't get list of courses from database, error: %w", result.Error)
	}
	return courses, nil
}
