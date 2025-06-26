package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//AnalysisLLM representa a resposta estruturada da LLM;
type AnalysisLLM struct {
	Score            float64  `json:"score,omitempty"`
	PontosPositivos  []string `json:"pontos_positivos,omitempty"`
	PontosNegativos  []string `json:"pontos_negativos,omitempty"`
	Recomendacoes    []string `json:"recomendacoes,omitempty"`
	FeedbackGeral    string   `json:"feedback_geral,omitempty"`
}

//BuildPromptForAnalysis monta o prompt MCP com currículo e vaga
func BuildPromptForAnalysis(cvText, jobDescription string) string {
	prompt := `Analise o currículo abaixo em relação à descrição da vaga. Responda APENAS em JSON com a seguinte estrutura:

{
	"score": <pontuação de 0 a 100>,
	"pontos_positivos": [ "texto", "texto" ],
	"pontos_negativos": [ "texto", "texto" ],
	"recomendacoes": [ "texto", "texto" ],
	"feedback_geral": "texto"
}

Currículo:
%s

Descrição da vaga:
%s`

	return fmt.Sprintf(prompt, cvText, jobDescription)
}

//SendPromptToLLM envia o prompt para a LLM local e retorna a resposta estruturada
func SendPromptToLLM(prompt string) (*AnalysisLLM, error) {
	bodyData := map[string]interface{}{
		"prompt": prompt,
		"model":  "IAResumeScanner",
		"stream": false,
		"options": map[string]interface{}{
			"temperature": 0.2,
		},
	}

	jsonData, err := json.Marshal(bodyData)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter dados para JSON: %w", err)
	}

	client := &http.Client{
		Timeout: time.Minute * 10,
	}

	req, err := http.NewRequest("POST", "http://localhost:11434/api/generate", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição HTTP: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição à LLM: %w", err)
	}
	defer response.Body.Close()

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta da LLM: %w", err)
	}

	var envelope struct {
		Response string `json:"response"`
	}
	if err := json.Unmarshal(responseBytes, &envelope); err != nil {
		return nil, fmt.Errorf("erro ao interpretar envelope da LLM: %w\nResposta bruta: %s", err, string(responseBytes))
	}

	var result AnalysisLLM
	if err := json.Unmarshal([]byte(envelope.Response), &result); err != nil {
		return nil, fmt.Errorf("erro ao interpretar JSON final: %w\nConteúdo: %s", err, envelope.Response)
	}

	return &result, nil
}