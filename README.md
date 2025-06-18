# 🧠 IA Resume Scanner

> Projeto para análise de currículos com IA generativa local (Ollama). Compare perfis com vagas, obtenha um score de compatibilidade e receba feedback inteligente com base em requisitos reais.

![Go](https://img.shields.io/badge/Golang-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white)
![Firebase](https://img.shields.io/badge/Firebase-FFCA28?style=for-the-badge&logo=firebase&logoColor=black)
![Ollama](https://img.shields.io/badge/Ollama-000000?style=for-the-badge&logo=llama&logoColor=white)
![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue?style=for-the-badge)
![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow?style=for-the-badge)

---

## 🎯 Objetivo

Criar uma plataforma inteligente para:
- 📥 Receber currículos (PDF, DOCX ou texto)
- 🧠 Analisar com IA generativa local (LLM via Ollama)
- 🆚 Comparar com os requisitos de uma vaga
- 🎯 Retornar um **score de compatibilidade**
- 🗣️ Fornecer **feedback em linguagem natural**

---

## ⚙️ Stack Utilizada

| Camada         | Tecnologia                                  |
|----------------|---------------------------------------------|
| **Backend**    | Go (Golang) + Gin ou Fiber                  |
| **IA**         | Ollama (modelos locais: Mistral, LLaMA etc.)|
| **Parser**     | Unidoc ou wrapper PDF-to-text em Go         |
| **Banco**      | PostgreSQL (Railway) e Firebase Firestore  |
| **Frontend**   | Next.js (React)                             |

---

## ☁️ Hospedagens

| Categoria    | Plataforma         | Destaques                                                        |
|--------------|--------------------|------------------------------------------------------------------|
| **Backend**  | Render, Railway, Fly.io | Deploy gratuito com integração GitHub                        |
| **LLM Local**| Oracle Cloud Free, WSL2, Docker+Ngrok | Rodar modelos Ollama localmente com até 24GB RAM        |
| **Banco**    | PostgreSQL (Railway), Firebase Firestore | Gratuitos e fáceis de integrar                             |
| **Frontend** | [Vercel](https://vercel.com) | Deploy ideal para projetos Next.js com CI/CD automáticos   |

---

## 🚀 Como Rodar Localmente

### ✅ Pré-requisitos

- [Go 1.21+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/) e [Firebase Firestore](https://firebase.google.com/)
- [Ollama](https://ollama.com/) instalado com modelo baixado (ex: `mistral`)
- [Node.js 18+](https://nodejs.org/) com Yarn ou NPM para o frontend (Next.js)

---

### 🖥️ Backend

```bash
git clone https://github.com/matheushermes/IAResumeScanner.git
cd IAResumeScanner
go mod tidy
cp .env.example .env
# Edite o .env com as configurações de:
# - Banco de dados (PostgreSQL ou Firebase)
# - Endpoint local do Ollama
go run main.g
```

### 🌐 Frontend (Next.js)
```bash
cd frontend
npm install
npm run dev
# - ou, se estiver usando Yarn:
yarn install
yarn dev
```