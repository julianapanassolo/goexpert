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

func buscarBrasilAPI(cep string) (*Endereco, string, error) {

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	client := &http.Client{Timeout: 1 * time.Second}
	resposta, err := client.Get(url)
	if err != nil {
		return nil, "brasilapi", err
	}
	defer resposta.Body.Close()

	if resposta.StatusCode != http.StatusOK {
		return nil, "brasilapi", fmt.Errorf("status code: %d", resposta.StatusCode)
	}

	body, err := ioutil.ReadAll(resposta.Body)
	if err != nil {
		return nil, "brasilapi", err
	}

	var endereco Endereco
	err = json.Unmarshal(body, &endereco)
	if err != nil {
		return nil, "brasilapi", err
	}

	return &endereco, "brasilapi", nil
}




func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	endereco, api, erro := buscaCEP(cep)
	if erro != nil {
		fmt.Println("erro:", erro)
		return
	}

	fmt.Printf("Endereço: %s, %s, %s - %s\n", endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.UF)
	fmt.Printf("API: %s\n", api)
}
