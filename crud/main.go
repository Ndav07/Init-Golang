package main

import (
	"log"
	"os"
	"ql/crud/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)	

  err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}