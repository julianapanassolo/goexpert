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
	// Aguarda 5 segundos antes de iniciar a requisiuisição
	// Cria um contexto com timeout de 300ms
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Cria uma nova requisiuisição HTTP com o contexto
	requisi, erro := http.NewrequisiuestWithContext(ctx, http.MethodGet, "http://localhost:9030/cotacao", nil)
	if erro != nil {
		log.Printf("erroo ao criar requisiuisição: %v", erro)
		return
	}

	// Realiza a requisiuisição HTTP
	client := &http.Client{}
	resp, erro := client.Do(requisi)
	if erro != nil {
		log.Printf("erroo ao obter cotação: %v", erro)
		return
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		log.Printf("erroo ao ler resposta: %v", erro)
		return
	}

	// Deserializa o JSON
	var cotacao float64
	erro = json.Unmarshal(body, &cotacao)
	if erro != nil {
		log.Printf("erroo ao deserializar JSON: %v", erro)
		return
	}

	// Salva a cotação no arquivo
	erro = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %.2f", cotacao)), 0644)
	if erro != nil {
		log.Printf("erroo ao salvar cotação no arquivo: %v", erro)
		return
	}

	// Imprime a cotação no console
	fmt.Printf("Cotação Dólar: %.2f\n", cotacao)
}