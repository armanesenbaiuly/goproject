package main

import (
	"github.com/armanesenbaiuly/go-rest-gorm/internal/db"
	"github.com/armanesenbaiuly/go-rest-gorm/internal/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	db.InitDB()
	r := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server started on: ", port)
	log.Fatal(r.Run(":" + port))
}
