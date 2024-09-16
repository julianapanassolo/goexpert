package main

import "database/sql"


type Cotacao struct {
	Bid float64 `json:"bid"`
}


func main() {
	db, error := sql.Open("sqlite3", "cotacoes.db")
	if erro 
}