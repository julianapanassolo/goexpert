package main

import (
	"requests"
	"threading"
	"time"

func main() {

	def busca_brasilapi(cep):
 		 url = f"https://brasilapi.com.br/api/cep/v1/{cep}"
  			try:
    		response = requests.get(url, timeout=1)
    		if response.status_code == 200:
      		return response.json(), "brasilapi"
    else:
      return None, "brasilapi"
  except requests.exceptions.Timeout:
    return None, "brasilapi"

def busca_viacep(cep):
  url = f"http://viacep.com.br/ws/{cep}/json/"
  try:
    response = requests.get(url, timeout=1)
    if response.status_code == 200:
      return response.json(), "viacep"
    else:
      return None, "viacep"
  except requests.exceptions.Timeout:
    return None, "viacep"

}