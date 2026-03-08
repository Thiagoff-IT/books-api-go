package model

import "strings"

// Livro representa a entidade de domínio de um livro.
type Livro struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Categoria string `json:"categoria"`
	Autor     string `json:"autor"`
	Sinopse   string `json:"sinopse"`
}

// Validar verifica se os campos obrigatórios estão preenchidos.
// Retorna uma lista de campos inválidos.
func (l *Livro) Validar() []string {
	var erros []string
	if strings.TrimSpace(l.Titulo) == "" {
		erros = append(erros, "titulo")
	}
	if strings.TrimSpace(l.Autor) == "" {
		erros = append(erros, "autor")
	}
	if strings.TrimSpace(l.Categoria) == "" {
		erros = append(erros, "categoria")
	}
	return erros
}
