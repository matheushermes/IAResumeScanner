package models

import (
	"errors"
	"strings"
)

//Job representa uma vaga de emprego;
type Job struct {
	JobDescription		string	`json:"description,omitempty"`
}

//Validations chama todos os métodos de validação;
func (j *Job) Validations() error {
	if err := j.checkJobDescription(); err != nil {
		return err
	}

	j.removeSpacesAtEnds()

	return nil
}

//checkJobDescription verifica se a descrição da vaga está valida;
func  (j *Job) checkJobDescription() error {
	if j.JobDescription == "" {
		return errors.New("A descrição da vaga ausente ou inválida")
	}

	return nil
}

//removeSpacesAtEnds remove os espaços em branco das extremidades;
func (j *Job) removeSpacesAtEnds() {
	j.JobDescription = strings.TrimSpace(j.JobDescription)
}