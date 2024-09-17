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

func buscaBrasilAPI(cep string) (*Endereco, string, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	cliente := &http.cliente{Timeout: 1 * time.Second}
	resp, err := cliente.Get(url)
	if err != nil {
		return nil, "brasilapi", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "brasilapi", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
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
