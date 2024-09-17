package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Cria uma nova requisição HTTP com o contexto
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9030/cotacao", nil)
	if err != nil {
		log.Printf("Erro ao criar requisição: %v", err)
		return
	}

	// Realiza a requisição HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Erro ao obter cotação: %v", err)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler resposta: %v", err)
		return
	}

	// Deserializa o JSON
	var cotacao float64
	err = json.Unmarshal(body, &cotacao)
	if err != nil {
		log.Printf("Erro ao deserializar JSON: %v", err)
		return
	}

	// Salva a cotação no arquivo
	err = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %.2f", cotacao)), 0644)
	if err != nil {
		log.Printf("Erro ao salvar cotação no arquivo: %v", err)
		return
	}

	// Imprime a cotação no console
	fmt.Printf("Cotação Dólar: %.2f\n", cotacao)
}