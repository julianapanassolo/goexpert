package main

import (
	"database/sql"
	"log"
)


type Cotacao struct {
	Bid float64 `json:"bid"`
}


func main() {
	db, error := sql.Open("sqlite3", "cotacoes.db")
	if erro := nil {
		log.Fatal(erro)
}
	defer db.Close()
