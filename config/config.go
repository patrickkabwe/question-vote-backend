package config

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "manage-budget-api"

func GetEnv(key string) string {

	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `.env`)

	currentEnv := os.Getenv("GO_ENV")

	log.Println("Current env: ", currentEnv)

	if currentEnv == "test" {
		return os.Getenv(key)
	}

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Return the value of the variable
	return os.Getenv(key)
}
