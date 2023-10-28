package main

import (
	"log"
	"os"
	"vote-app/database"
	"vote-app/internal/routes"
	"vote-app/utils"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	app := utils.CreateApp()
	db, err := database.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	routes.SetupRoutes(db, app)
	log.Fatal(app.Listen("0.0.0.0:" + port))
}