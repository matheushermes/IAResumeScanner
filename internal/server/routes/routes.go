package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/IAResumeScanner/pkg/controllers"
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

	return router
}