package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func notfound(c *gin.Context) {
	c.String(404, "404 - Service Not Found")
}

//InitialiseRouter - Initialise the Router and associated Route Handlers
func InitialiseRouter() {
	//Intialise Router
	router := gin.New()

	//Middlewares
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.NoRoute(notfound)

	//Intialise Router Group
	v1 := router.Group("/")
	{
		v1.GET("/", notfound)
	}

	//Execute
	router.Run(":8080")
}
