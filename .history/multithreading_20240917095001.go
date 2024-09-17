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

func main() {
	var cep string
	fmt.Print("Digite o  CEP: ")
	fmt.Scanln(&cep)

	endereco, api, erro := buscarCEP(cep)
	if erro != nil {
		fmt.Println("erro:", erro)
		return
	}

	fmt.Printf("Endereço: %s, %s, %s - %s\n", endereco.Logradouro, endereco.Bairro, endereco.Localidade, endereco.UF)
	fmt.Printf("API: %s\n", api)

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
	ch := make(chan struct {
		endereco *Endereco
		api      string
		erro     error
	}, 2)

	// Inicia as threads para buscar na BrasilAPI e ViaCEP
	go func() {
		defer wg.Done()
		endereco, api, erro := buscarBrasilAPI(cep)
		ch <- struct {
			endereco *Endereco
			api      string
			erro     error
		}{endereco, api, erro}
	}()

	go func() {
		defer wg.Done()
		endereco, api, erro := buscarViaCEP(cep)
		ch <- struct {
			endereco *Endereco
			api      string
			erro     error
		}{endereco, api, erro}
	}()

	// Aguarda o término das threads
		wg.Wait()
		close(ch)


	// Seleciona o resultado mais rápido
	for i := 0; i < 2; i++ {
		select {
		case result := <-ch:
			if result.erro == nil {
				resultado = result.endereco
				api = result.api
				return resultado, api, nil
			}
		default:
			// Se nenhuma resposta chegou, é timeout em ambas
			erro = fmt.Errorf("timeout em ambas as APIs")
		}
	}

	return resultado, api, erro
}






