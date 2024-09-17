package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Representa: O endereço retornado pelas APIs
type Endereco struct {
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}


// Função: Realiza a requisição à API BrasilAPI
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


// Função: Realiza a requisição à API ViaCEP
func buscarViaCEP(cep string) (*Endereco, string, error) {

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	client := &http.Client{Timeout: 1 * time.Second}
	resposta, erro := client.Get(url)
	if erro != nil {
		return nil, "viacep", erro
	}
	defer resposta.Body.Close()

	if resposta.StatusCode != http.StatusOK {
		return nil, "viacep", fmt.Errorf("status code: %d", resposta.StatusCode)
	}

	body, erro := ioutil.ReadAll(resposta.Body)
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


// Função: Gerência as threads e buscar o endereço
func buscarCEP(cep string) (*Endereco, string, error) {
	var wg sync.WaitGroup
	wg.Add(2)

	var resultado *Endereco
	var api string
	var erro error

	// Cria um canal para comunicação entre as threads
	ch := make(chan *Endereco, 2)

	// Inicia as threads para buscar na BrasilAPI e ViaCEP
	go func() {
		defer wg.Done()
		endereco, api, erro := buscarBrasilAPI(cep)
		if erro == nil {
			ch <- endereco
		}
	}()

	go func() {
		defer wg.Done()
		endereco, api, erro := buscarViaCEP(cep)
		if erro == nil {
			ch <- endereco
		}
	}()

	// Aguarda o término das threads
	wg.Wait()
	close(ch)

	// Seleciona o resultado mais rápido
	select {
	case endereco := <-ch:
		resultado = endereco
		api = "brasilapi"
	case endereco := <-ch:
		resultado = endereco
		api = "viacep"
	default:
		erro = fmt.Errorf("timeout em ambas as APIs")
	}

	return resultado, api, err
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
