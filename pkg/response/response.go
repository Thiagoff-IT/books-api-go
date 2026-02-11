package response

import (
	"encoding/json"
	"net/http"
)

// JSON escreve uma resposta JSON com o status code informado.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// Erro escreve uma resposta de erro padronizada em JSON.
func Erro(w http.ResponseWriter, statusCode int, mensagem string) {
	JSON(w, statusCode, map[string]string{"erro": mensagem})
}
