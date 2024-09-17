package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)


type Endereco struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	endereco, api, erro := buscaCEP(cep)
	if erro != nil {
		fmt.Println("erroo:", erro)
		return
	}

	fmt.Printf("Endereço: %s, %s, %s - %s\n", endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.UF)
	fmt.Printf("API: %s\n", api)
}
