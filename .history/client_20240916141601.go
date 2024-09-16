package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// Define o contexto com timeout para a requisição.
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Faz a requisição ao server.go.
	resp, err := http.GetWithContext(ctx, "http://localhost:8080/cotacao")
	if err != nil {
		log.Printf("Erro ao obter cotação: %v", err)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta: %v", err)
		return
	}

	// Deserializa o JSON.
	var cotacao float64
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Printf("Erro ao deserializar JSON: %v", err)
		return
	}

	// Salva a cotação no arquivo cotacao.txt.
	err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %.2f", cotacao)), 0644)
	if err != nil {
		log.Printf("Erro ao salvar cotação no arquivo: %v", err)
		return
	}

	fmt.Printf("Cotação do Dólar: %.2f\n", cotacao)
}