package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/IAResumeScanner/pkg/controllers"

	
	_ "github.com/matheushermes/IAResumeScanner/api"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRouter(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		scanner := main.Group("scanner")
		{
			scanner.POST("/upload", controllers.UploadFile)
			scanner.POST("/match", controllers.MatchCV)
		}
	}

	//Rota Swagger;
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}