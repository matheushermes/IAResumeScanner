# 🧠 IA Resume Scanner

Projeto para análise de currículos com IA generativa local (Ollama). Compare perfis com vagas, obtenha um score de compatibilidade e receba feedback inteligente com base em requisitos reais.

![Go](https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Ollama](https://img.shields.io/badge/Ollama-000000?style=for-the-badge&logo=llama&logoColor=white)
![UniDoc](https://img.shields.io/badge/UniDoc-PDF%2FDOCX-blue?style=for-the-badge)
![Swagger](https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black)
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

---
  
# 🧠 Criar modelo customizado no Ollama

Antes de rodar o projeto, você deve criar um modelo personalizado com o nome IAResumeScanner usando o mistral como base:

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
---

# 🚀 Como Rodar Localmente

## 1. Clone o projeto

```bash
git clone https://github.com/matheushermes/IAResumeScanner.git
cd IAResumeScanner
```
## 2. Instale as dependências

Antes de iniciar o projeto, certifique-se de instalar todas as dependências necessárias executando o comando abaixo:

```bash
go mod tidy
```

## 3. Configure o ambiente

Copie o arquivo de exemplo .env.example para criar seu próprio arquivo de configuração .env:
```bash
cp .env.example .env
```

Em seguida, edite o arquivo .env com os valores apropriados para o seu ambiente:
```env
API_PORT=5000
UNIDOC_LICENSE_API_KEY=SUACHAVEAQUI
```

## 4. Inicie o modelo no Ollama

Para iniciar o modelo localmente no Ollama, execute o comando abaixo no terminal:
```bash
ollama run IAResumeScanner
```
Esse comando inicializa o modelo IAResumeScanner previamente configurado no Ollama, permitindo que ele fique disponível para receber requisições na sua máquina local.

> ⚠️ Certifique-se de que o Ollama está instalado e que o modelo IAResumeScanner foi criado corretamente antes de executar este comando.

## 5. Rode o servidor

Para iniciar o backend, execute o comando abaixo no terminal:
```bash
go run cmd/main.go
```
O servidor estará disponível em:
📍 http://localhost:5000

---

# 📂 Endpoints da API
### ➕ Upload de currículo
```http
POST /api/v1/scanner/upload
```
- <b>Descrição:</b> Endpoint para envio do arquivo do currículo.
- <b>Formato do arquivo:</b> `multipart/form-data` contendo o campo `file` com extensão `.pdf` ou `.docx`.

### 📊 Análise de Compatibilidade com a Vaga
```http
POST /api/v1/scanner/match
```
- <b>Descrição:</b> Recebe a descrição completa da vaga e retorna uma avaliação da compatibilidade com o currículo enviado anteriormente.

<b><i>Corpo da Requisição (JSON):</b></i>
```json
{
  "jobDescription": "Descrição completa da vaga"
}
```

<b><i>Resposta esperada:</b></i>
```json
{
  "score": 85,
  "pontos_positivos": ["texto", "texto"],
  "pontos_negativos": ["texto", "texto"],
  "recomendacoes": ["texto", "texto"],
  "feedback_geral": "texto"
}
```
---

# 📄 Documentação Swagger
Após iniciar o servidor, acesse a documentação interativa da API no navegador:
```http
http://localhost:5000/swagger/index.html
```
---

# 📌 Observações Técnicas

- A análise utiliza prompt estruturado via MCP (Most Common Prompting), garantindo respostas consistentes e formatadas em JSON.
- O parser aceita arquivos nos formatos `.pdf` e `.docx`, extraindo o conteúdo através da biblioteca UniDoc.
- A comunicação com a LLM é feita via chamadas HTTP locais ao serviço Ollama, com timeout configurado para até 10 minutos, suportando análises extensas.

# 📄 Licença
Este projeto está licenciado sob a <b>MIT License</b>.

<p align="center"> <strong>🚀 Desenvolvido por Matheus Hermes com foco em IA prática, local e ética.</strong> </p>