package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/matheushermes/IAResumeScanner/pkg/models"
	"github.com/matheushermes/IAResumeScanner/pkg/utils"
)

//GetFirstFileFromUploads retorna o caminho do primeiro arquivo dentro da pasta ../docs/uploads;
func GetFirstFileFromUploads() (string, error) {
	files, err := os.ReadDir("../docs/uploads")
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() {
			return filepath.Join("../docs/uploads", file.Name()), nil
		}
	}

	return "", fmt.Errorf("nenhum arquivo encontrado na pasta uploads")
}

//MatchCV compara o conteúdo de um currículo com os requisitos de uma vaga, retornando uma pontuação e compatibilidade do candidato com a vaga;
// @Summary Analisa o currículo em relação à vaga
// @Description Faz análise de compatibilidade entre currículo enviado e a descrição da vaga, usando LLM local
// @Tags Scanner
// @Accept json
// @Produce json
// @Param job body models.Job true "Descrição da Vaga"
// @Success 200 {object} utils.AnalysisLLM
// @Failure 400 {object} map[string]string
// @Router /api/v1/scanner/match [post]
func MatchCV(c *gin.Context) {
	filePath, err := GetFirstFileFromUploads()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao obter o arquivo do CV: " + err.Error(),
		})
		return
	}

	cvText, err := utils.ExtractTextFromCV(filePath)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Erro ao extrair texto do CV: " + err.Error(),
		})
		return
	}

	var jobDescription models.Job
	if err := c.ShouldBindJSON(&jobDescription); err != nil {
		c.JSON(422, gin.H {
			"error": err.Error(),
		})
		return
	}

	prompt := utils.BuildPromptForAnalysis(cvText, jobDescription.JobDescription)

	resultAnalysisLLM, err := utils.SendPromptToLLM(prompt)
	if err != nil {
		c.JSON(500, gin.H {
			"error": "Erro ao interpretar JSON da LLM",
			"rawText": err.Error(),
		})
		return
	}

	c.JSON(200, resultAnalysisLLM)
}