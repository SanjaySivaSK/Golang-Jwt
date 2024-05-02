package main

import (
	"github.com/gin-gonic/gin"

	"MyModulw/controller"
	"MyModulw/Database"
	"MyModulw/middleware"
)

func main() {
	
	database.Connect("root:root@tcp(localhost:5432)/coco2?parseTime=true")
	database.Migrate()

	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
	}
	return router


}
