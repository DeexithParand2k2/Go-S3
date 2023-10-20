package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	"Go-S3/internal/routes"
	"github.com/joho/godotenv"
	"log"
)

func StartServer(){
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default();

	if godotenv.Load()!=nil {
		log.Fatalf("Error loading in .env file : ",godotenv.Load())
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	addr := host+":"+port

	//single test route - greeting - router.group
	singletestroute.SetupSingleTestRoute(router)

	fmt.Println("Server running at localhost",addr)


	router.Run(addr)
}