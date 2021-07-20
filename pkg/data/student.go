package data

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
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

func (d StudentData) Add(student Student) (int, error) {
	result := d.db.Create(&student)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create student, error: %w", result.Error)
	}
	return student.Id, nil
}

func (d StudentData) Read(id int) (Student, error) {
	var student Student
	result := d.db.Find(&student, id)
	if result.Error != nil {
		return student, fmt.Errorf("can't read student with given id, error: %w", result.Error)
	}
	return student, nil
}

func (d StudentData) ReadAll() ([]Student, error) {
	var students []Student
	result := d.db.Find(&students)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read students from database, error: %w", result.Error)
	}
	return students, nil
}

func (d StudentData) Delete(id int) error {
	result := d.db.Delete(&Student{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't delete student from database, error: %w", result.Error)
	}
	return nil
}

func notePayment(tx *gorm.DB, studentId, courseId int, passed bool) error {
	err := tx.Create(&Payment{
		StudentId: studentId,
		CourseId:  courseId,
		Date:      time.Now().Format("2006-01-02 15:04:05"),
		Passed:    passed,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (d StudentData) Pay(studentId, courseId, payment int) error {
	tx := d.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return fmt.Errorf("can't start a transaction, err: %w", err)
	}

	var course Course
	if err := tx.Model(&Course{}).Where("id = ?", courseId).Find(&course).Error; err != nil {
		tx.Rollback()
		err = notePayment(d.db.Begin(), studentId, courseId, false)
		if err != nil {
			log.Errorf("can't create payment, err: %v", err)
		}
		return fmt.Errorf("can't find course with given id, err: %w", err)
	}

	if payment < course.Fee {
		tx.Rollback()
		err := notePayment(d.db.Begin(), studentId, courseId, false)
		if err != nil {
			log.Errorf("can't create payment, err: %v", err)
		}
		return fmt.Errorf("insufficient funds")
	}

	if err := tx.Model(&Student{}).Where("id = ?", studentId).Update("is_active", true).Error; err != nil {
		tx.Rollback()
		err = notePayment(d.db.Begin(), studentId, courseId, false)
		if err != nil {
			log.Errorf("can't create payment, err: %v", err)
		}
		return fmt.Errorf("can't update student status, err: %w", err)
	}

	return notePayment(tx, studentId, courseId, true)
}
