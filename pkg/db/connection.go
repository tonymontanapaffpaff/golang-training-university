package db

import (
	"fmt"

	"github.com/tonymontanapaffpaff/golang-training-university/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(config config.Configuration) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host, config.Port, config.User, config.DBName, config.Password, config.SSLMode)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("got an error when tried to make connection with database:%w", err)
	}

	return connection, nil
}
