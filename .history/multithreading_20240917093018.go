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
	resposta, erro := client.Get(url)
	if erro != nil {
		return nil, "brasilapi", erro
	}
	defer resposta.Body.Close()

	if resposta.StatusCode != http.StatusOK {
		return nil, "brasilapi", fmt.Errorf("status code: %d", resposta.StatusCode)
	}

	body, erro := ioutil.ReadAll(resposta.Body)
	if erro != nil {
		return nil, "brasilapi", erro
	}

	var endereco Endereco
	erro = json.Unmarshal(body, &endereco)
	if erro != nil {
		return nil, "brasilapi", erro
	}

	return &endereco, "brasilapi", nil
}


func buscarViaCEP(cep string) (*Endereco, string, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	client := &http.Client{Timeout: 1 * time.Second}
	resposta, erro := client.Get(url)
	if erro != nil {
		return nil, "viacep", err
	}
	defer resposta.Body.Close()

	if resposta.StatusCode != http.StatusOK {
		return nil, "viacep", fmt.Errorf("status code: %d", resposta.StatusCode)
	}

	body, erro := ioutil.ReadAll(resp.Body)
	if erro != nil {
		return nil, "viacep", erro
	}

	var endereco Endereco
	erro = json.Unmarshal(body, &endereco)
	if erro != nil {
		return nil, "viacep", erro
	}

	return &endereco, "viacep", nil
}




func main() {
	var cep string
	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	endereco, api, erro := buscarCEP(cep)
	if erro != nil {
		fmt.Println("erro:", erro)
		return
	}

	fmt.Printf("Endereço: %s, %s, %s - %s\n", endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.UF)
	fmt.Printf("API: %s\n", api)
}
