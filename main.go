package main

import (
	"fmt"
	"Go-S3/internal/server"
)

func main(){
	server.StartServer();
	fmt.Println("End of main file")
}