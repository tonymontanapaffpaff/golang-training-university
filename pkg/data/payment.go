package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Payment struct {
	Id        int    `gorm:"primaryKey"`
	StudentId int    `gorm:"student_id"`
	CourseId  int    `gorm:"course_id"`
	Date      string `gorm:"date"`
	Passed    bool   `gorm:"passed"`
}

type PaymentData struct {
	db *gorm.DB
}

func NewPaymentData(db *gorm.DB) *PaymentData {
	return &PaymentData{db: db}
}

func (d PaymentData) ReadAll() ([]Payment, error) {
	var payments []Payment
	result := d.db.Find(&payments)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read payments from database, error: %w", result.Error)
	}
	return payments, nil
}
