package controllers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Arquivo não encontrado",
		})
		return
	}

	filename := filepath.Base(fileHeader.Filename)

	uploadDir, err := filepath.Abs("../docs/uploads")
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao resolver diretório de upload",
		})
		return
	}

	filePath := filepath.Join(uploadDir, filename)

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao criar diretório de upload",
		})
		return
	}

	dst, err := os.Create(filePath)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao criar arquivo para salvar",
		})
		return
	}
	defer dst.Close()

	src, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao abrir o arquivo",
		})
		return
	}
	defer src.Close()

	if _, err := io.Copy(dst, src); err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao copiar conteúdo do arquivo",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Arquivo armazenado com sucesso!",
	})
}