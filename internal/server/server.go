package server

import (
	"github.com/gin-gonic/gin"
	"os"
	"fmt"
	//"Go-S3/internal/handlers"
)

func StartServer(){
	router := gin.Default();

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	addr := host+":"+port

	fmt.Println("Server running at localhost",addr)

	//router.GET("/",)

	router.Run(addr)
}