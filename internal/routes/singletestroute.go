package singletestroute

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func SetupSingleTestRoute(router *gin.Engine){
	SingleTestRoute := router.Group("/greeting")
	{
		SingleTestRoute.GET("/sayhello",SayHello)
		SingleTestRoute.GET("/saygoodbye",SayGoodBye)
	}
}

func SayHello(c *gin.Context){

	fmt.Println("Get Request -> hello boyz")

	c.IndentedJSON(http.StatusOK,gin.H{"message": "Hello Boyz"})
}

func SayGoodBye(c *gin.Context){

	fmt.Println("Get Request -> Goodbye boyz")

	c.IndentedJSON(http.StatusOK,gin.H{"message": "Good Bye Boyzz"})
}