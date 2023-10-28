package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

func Connect() (*DB, error) {
	dbUrl := "postgres://postgres:postgres@localhost:5432/vote?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	return &DB{
		DB: db,
	}, err
}

func (db *DB) Close() {

}
