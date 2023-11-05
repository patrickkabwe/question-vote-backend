package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"vote-app/internal/models"
)

type Database struct {
	*gorm.DB
}

func Connect(dbUrl string) (*Database, error) {
	_db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return &Database{
		DB: _db,
	}, err
}

func (db *Database) Migrate() error {
	err := db.AutoMigrate(
		&models.Member{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() {
	sqlDB, _ := db.DB.DB()
	err := sqlDB.Close()
	if err != nil {
		return
	}
}
