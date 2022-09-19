package main

import (
	_ "gin_swagger/docs"
	"gin_swagger/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

/*=== Swagger API Document Title ===*/
// @title     Gin Swagger API Document
// @version         1.0
// @description     A book management service API in Go using Gin framework.
// @host      localhost:8080
// @BasePath  /api/v1
/*=== Swagger API Document Title ===*/

func main() {
	router := routes.SetupRouter()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
