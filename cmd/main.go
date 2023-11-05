package main

import (
	"log"
	"vote-app/config"
	"vote-app/database"
	"vote-app/internal/routes"
	"vote-app/utils"
)

func main() {
	port := config.GetEnv("PORT")
	dbUrl := config.GetEnv("DATABASE_URL")
	if port == "" {
		port = "5000"
	}
	if dbUrl == "" {
		log.Fatalln("DATABASE_URL is not set")
	}
	app := utils.CreateApp()

	db, err := database.Connect(dbUrl)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Migrate()
	if err != nil {
		log.Fatalln(err)
	}

	routes.SetupRoutes(db, app)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
