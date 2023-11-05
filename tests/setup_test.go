package tests

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"testing"
	"vote-app/database"
)

var gdb *gorm.DB

func setupDB() *database.Database {
	return &database.Database{
		DB: gdb,
	}
}

func TestMain(m *testing.M) {
	// setup
	fmt.Println("Before tests")
	db, err := database.Connect("postgres://postgres:postgres@localhost:5432/vote_app_test?sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Migrate()
	if err != nil {
		fmt.Println(err)
		return
	}

	gdb = db.DB

	// run tests
	runner := m.Run()

	allTables, err := db.Migrator().GetTables()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, table := range allTables {
		err = db.Migrator().DropTable(table)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// teardown
	db.Close()

	// exit
	os.Exit(runner)
}
