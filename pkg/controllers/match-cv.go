package controllers

import "github.com/gin-gonic/gin"

func MatchCV(c *gin.Context) {
	c.JSON(200, gin.H {
		"message": "sucesso!",
	})
}