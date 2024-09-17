// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func main() {

// 	ctx, cancelado := context.WithTimeout(context.Background(), 300*time.Millisecond)
// 	defer cancelado()

// 	resposta, erro := http.GetWithContext(ctx, "http://localhost:9030/cotacao")
// 	if erro != nil {
// 		log.Printf("Erro ao obter cotação: %v", erro)
// 		return
// 	}
// 	defer resposta.Body.Close()

// 	body, erro := ioutil.ReadAll(resposta.Body)
// 	if erro != nil {
// 		log.Printf("Erro ao ler resposta: %v", erro)
// 		return
// 	}

// 	var cotacao float64
// 	erro = json.Unmarshal(body, &cotacao)
// 	if erro != nil {
// 		log.Printf("Erro ao deserializar JSON: %v", erro)
// 		return
// 	}

// 	erro = ioutil.WriteFile("cotacao.txt", []byte(fmt.Sprintf("Dólar: %.2f", cotacao)), 0644)
// 	if erro != nil {
// 		log.Printf("Erro ao salvar cotação no arquivo: %v", erro)
// 		return
// 	}

// 	fmt.Printf("Cotação Dólar: %.2f\n", cotacao)
// }
