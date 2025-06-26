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
	prompt := `Você é um assistente de análise de compatibilidade entre currículo e vaga. Seu objetivo é ajudar um candidato a entender o quanto seu perfil atual atende aos requisitos de uma vaga específica.

Abaixo estão dois blocos de texto:
- O primeiro é o CURRÍCULO do candidato.
- O segundo é a DESCRIÇÃO DA VAGA desejada.

Sua tarefa é analisar **somente as informações disponíveis** e gerar um feedback objetivo e útil para o candidato. A resposta deve ser **EXCLUSIVAMENTE** em JSON válido no seguinte formato:

{
  "score": <número entre 0 e 100 representando o grau de compatibilidade>,
  "pontos_positivos": [ "fatores do currículo que alinham bem com a vaga" ],
  "pontos_negativos": [ "fatores do currículo que não atendem ou estão ausentes em relação à vaga" ],
  "recomendacoes": [ "áreas ou competências que o candidato deve desenvolver para melhorar o match com a vaga" ],
  "feedback_geral": "resumo objetivo da análise com tom encorajador e informativo, voltado diretamente ao candidato"
}

⚠️ Responda SOMENTE com o JSON descrito acima, sem comentários ou explicações extras.

CURRÍCULO DO CANDIDATO:
%s

DESCRIÇÃO DA VAGA:
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