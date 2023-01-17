package main

import (
	"log"
	"os"
	"ql/crud/src/configuration/logger"
	"ql/crud/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user app")
	godotenv.Load()
	port := os.Getenv("PORT")
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)	

  err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}