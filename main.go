package main

import (
	urls "github.com/cesarcruzc/meli/api"
	"github.com/cesarcruzc/meli/config"
	_ "github.com/cesarcruzc/meli/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

// @title Meli Microservice API
// @version 1.0
// @description This is a microservice that consumes data from Mercado Libre and stores it in a database.

// @host 127.0.0.1:8888
// @BasePath /
func main() {
	// Inicializamos el logger
	logger := config.InitLogger()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inicializamos la conexión a la base de datos
	database, err := config.GetMongoClient()
	if err != nil {
		logger.Fatalf("Error getting mongo client: %s", err)
	}

	// Inicializamos el servidor
	gin.SetMode(gin.ReleaseMode)

	// Cargamos las rutas
	server := urls.LoadUrls(database, logger)

	// Corremos el servidor
	logger.Println("Server running on", ":8888")
	logger.Fatalf("Error in server: %s", server.Run(":8888"))
}
