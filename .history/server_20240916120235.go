package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	_ "github.com/mattn/go-sqlite3"
)


type Cotacao struct {
	Bid float64 `json:"bid"`
}


func main() {
	db, erro := sql.Open("sqlite3", "cotacoes.db")
	if error != nil {
		log.Fatal(error)
	}	
	defer db.Close()
	


_, erro = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, data TEXT, cotacao REAL)")
if erro != nil {
	log.Fatal(erro)
}

http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request){
	ctx, cancelar := context.WithTimeOut(context.Background(), 200*time.Millisecond)
	defer cancelar()

	resposta, erro := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if erro != nil {
		log.Printf("Erro ao fazer requisição: %v", erro)
		http.Error (w, "Erro ao obter cotação", http.StatusInternalServerError)
		return
	} 
	defer resposta.Body.Close()

	body, erro := ioutil.ReadAll(resposta.Body)
	if erro != nil {
		log.Printf("Erro ao ler resposta: %v", erro)
		http.Error (w, "Erro ao obter cotação", http.StatusInternalServerError)
		return
	}

	var cotacao Cotacao
	erro = json.Unmarshal(body, &cotacao)
	if erro != nil {
		log.Printf("Erro ao decodificar resposta: %v", erro)
		http.Error (w, "Erro ao obter cotação", http.StatusInternalServerError)
		return
	}

	ctx, cancelar := context.WithTimeOut(context.Background(), 10*time.Millisecond)
	defer cancelar()

	_, erro = db.ExecContext(ctx, "INSERT INTO cotacoes (data, cotacao) VALUES (?, ?)", time.Now(), cotacao.Bid)
	if erro != nil {
		log.Printf("Erro ao inserir cotação: %v", erro)
		http.Error (w, "Erro ao obter cotação", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cotacao.Bid)
})

fmt.Printf("Servidor iniciado na porta 9030\n")
	log.Fatal(http.ListenAndServe(":9030", nil))