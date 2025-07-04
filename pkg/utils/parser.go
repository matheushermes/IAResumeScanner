package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/unidoc/unioffice/v2/document"
	"github.com/unidoc/unipdf/v4/extractor"
	"github.com/unidoc/unipdf/v4/model"
)

//ExtractTextFromCV identifica a extensão do arquivo de currículo e delega a extração de texto para a função correspondente;
func ExtractTextFromCV(filePath string) (string, error) {
	ext := strings.ToLower(filepath.Ext(filePath))

	switch ext {
	case ".pdf":
		return extractTextFromPDF(filePath)
	case ".docx":
		return extractTextFromDOCX(filePath)
	default:
		return "", fmt.Errorf("formato não suportado: %s", ext)
	}
}

//extractTextFromPDF abre um arquivo PDF e extrai todo o conteúdo textual, página por página, utilizando a biblioteca UniPDF;
func extractTextFromPDF(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir o arquivo PDF: %w", err)
	}
	defer f.Close()

	reader, err := model.NewPdfReader(f)
	if err != nil {
		return "", fmt.Errorf("erro ao criar o leitor PDF: %w", err)
	}

	isEncrypted, err := reader.IsEncrypted()
	if err != nil {
		return "", fmt.Errorf("erro ao verificar criptografia do PDF: %w", err)
	}
	if isEncrypted {
		ok, err := reader.Decrypt([]byte(""))
		if err != nil {
			return "", fmt.Errorf("erro ao descriptografar PDF: %w", err)
		}
		if !ok {
			return "", fmt.Errorf("falha ao descriptografar PDF: senha incorreta ou não suportado")
		}
	}

	numPages, err := reader.GetNumPages()
	if err != nil {
		return "", fmt.Errorf("erro ao obter número de páginas: %w", err)
	}

	var sb strings.Builder
	for i := 1; i <= numPages; i++ {
		page, err := reader.GetPage(i)
		if err != nil {
			return "", fmt.Errorf("erro ao obter página %d: %w", i, err)
		}
		ex, err := extractor.New(page)
		if err != nil {
			return "", fmt.Errorf("erro ao criar extractor: %w", err)
		}
		text, err := ex.ExtractText()
		if err != nil {
			return "", fmt.Errorf("erro ao extrair texto da página %d: %w", i, err)
		}

		sb.WriteString(text + "\n")
	}

	return sb.String(), nil
}

//extractTextFromDOCX abre um arquivo .docx (Word) e extrai o texto de todos os parágrafos, concatenando o conteúdo de cada run (bloco de texto com formatação);
func extractTextFromDOCX(filePath string) (string, error) {
	doc, err := document.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir .docx: %w", err)
	}

	var sb strings.Builder
	for _, para := range doc.Paragraphs() {
        var paraText strings.Builder
        for _, run := range para.Runs() {
            paraText.WriteString(run.Text())
        }
        text := strings.TrimSpace(paraText.String())
        if text != "" {
            sb.WriteString(text + "\n")
        }
    }

    return sb.String(), nil
}