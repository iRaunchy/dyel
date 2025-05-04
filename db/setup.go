package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Init opens a connection using the provided DSN and runs AutoMigrate.
func Init(dsn string) (*gorm.DB, error) {
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("couldn't connect to database: %w", err)
	}
	// auto-migrate your models
	if err := dbConn.AutoMigrate(&Program{}, &Day{}, &Exercise{}); err != nil {
		return nil, fmt.Errorf("auto-migrate failed: %w", err)
	}

	return dbConn, nil

}
