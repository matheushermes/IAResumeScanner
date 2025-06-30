package controllers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//UploadFile recebe um arquivo enviado pelo usuário;
// @Summary Upload de currículo
// @Description Faz upload de um arquivo PDF ou DOCX contendo o currículo do candidato
// @Tags Scanner
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Arquivo do currículo (.pdf ou .docx)"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/scanner/upload [post]
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