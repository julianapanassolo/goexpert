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
}	
defer db.Close()

_, erro = db.Exec("CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, data TEXT, cotacao REAL)")
if erro := nil {
	log.Fatal(erro)
}

http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request){
	ctx, cancelar := context.WithTimeOut(context.Backg)
})
