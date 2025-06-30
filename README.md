# 🧠 IA Resume Scanner

> Projeto para análise de currículos com IA generativa local (Ollama). Compare perfis com vagas, obtenha um score de compatibilidade e receba feedback inteligente com base em requisitos reais.

![Go](https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Ollama](https://img.shields.io/badge/Ollama-000000?style=for-the-badge&logo=llama&logoColor=white)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)
![UniDoc](https://img.shields.io/badge/UniDoc-PDF%2FDOCX-blue?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow?style=for-the-badge)

---

## 🎯 Objetivo

Criar uma plataforma inteligente para:
- 📥 Receber currículos (PDF ou DOCX)
- 🧠 Analisar com IA generativa local (LLM via Ollama)
- 🆚 Comparar com os requisitos de uma vaga
- 🎯 Retornar um **score de compatibilidade**
- 🗣️ Fornecer **feedback técnico personalizado**

---

## ⚙️ Stack Utilizada

| Camada        | Tecnologia                                  |
|---------------|----------------------------------------------|
| **Backend**   | Go (Golang) + Gin                           |
| **IA Local**  | Ollama + modelo `mistral` customizado       |
| **Parser**    | UniPDF / UniOffice (via UniDoc)             |
| **Docs API**  | Swaggo + Swagger UI                         |

---

## ☁️ Requisitos Locais

- Go 1.20 ou superior
- [Ollama](https://ollama.com) instalado localmente
- Modelo LLM criado com prompt customizado (veja abaixo)
- Chave UniDoc válida (para extrair texto de `.pdf` e `.docx`)
- Swag instalado globalmente (para Swagger):
  
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```
# 🧠 Criar modelo customizado no Ollama

> Antes de rodar o projeto, você deve criar um modelo personalizado com o nome IAResumeScanner usando o mistral como base:

```bash
curl -X POST http://localhost:11434/api/create \
  -H "Content-Type: application/json" \
  -d '{
    "model": "IAResumeScanner",
    "from": "mistral",
    "system": "Você é um assistente de recrutamento especializado. Sua personalidade é muito profissional, precisa, objetiva e cordial. Você tem como responsabilidade analisar detalhadamente currículos e descrições de vagas, identificar habilidades, experiências, formações e lacunas, e fornecer feedback construtivo para ajudar candidatos a melhorarem seu perfil. Você deve focar em clareza, objetividade e insights úteis para seleção e desenvolvimento profissional.",
    "parameters": {
      "temperature": 0.3,
      "top_p": 0.9
    }
  }'
```

# 🚀 Como Rodar Localmente

## 1. Clone o projeto

```bash
git clone https://github.com/matheushermes/IAResumeScanner.git
cd IAResumeScanner
```
## 2. Instale as dependências

```bash
go mod tidy
```
## 3. Configure o ambiente

```bash
cp .env.example .env
```
Edite o arquivo `.env` com as seguintes variáveis:
```env
API_PORT=5000
UNIDOC_LICENSE_API_KEY=SUACHAVEAQUI
```
## 4. Gere a documentação Swagger

```bash
swag init --generalInfo cmd/main.go --output api/docs
```

## 5. Inicie o modelo no Ollama

```bash
ollama run IAResumeScanner
```
## 6. Rode o servidor

```bash
go run cmd/main.go
```
O backend estará disponível em:
📍 http://localhost:5000

# 📂 Endpoints da API
## ➕ Upload de currículo

```http
POST /api/v1/scanner/upload
```
- Form-data: file (.pdf ou .docx)

## 📊 Match com a vaga
```http
POST /api/v1/scanner/match
```
Body JSON:
```json
{
  "jobDescription": "Descrição completa da vaga"
}
```

Resposta esperada:
```json
{
  "score": 85,
  "pontos_positivos": ["texto", "texto"],
  "pontos_negativos": ["texto", "texto"],
  "recomendacoes": ["texto", "texto"],
  "feedback_geral": "texto"
}
```

# 📄 Documentação Swagger
Acesse no navegador após iniciar o servidor:

```http
http://localhost:5000/swagger/index.html
```
# 📌 Observações Técnicas

- A análise é feita via prompt estruturado (MCP) garantindo consistência no JSON.
- O parser de currículos aceita .pdf e .docx, extraindo o conteúdo com UniDoc.
- A comunicação com a LLM é feita via HTTP local (Ollama) com timeout de até 10 minutos.

# 📄 Licença
Este projeto está licenciado sob a MIT License.

<p align="center"><strong>🚀 Desenvolvido por Matheus Hermes com foco em IA prática, local e ética.</strong></p>