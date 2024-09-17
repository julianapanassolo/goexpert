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


		def busca_cep(cep):
  		thread1 = threading.Thread(target=busca_brasilapi, args=(cep,))
  thread2 = threading.Thread(target=busca_viacep, args=(cep,))

  thread1.start()
  thread2.start()

  thread1.join(1)
  thread2.join(1)

  if thread1.is_alive():
    thread2.join()
  elif thread2.is_alive():
    thread1.join()

  result1 = thread1.get_result()
  result2 = thread2.get_result()

  if result1 is not None:
    return result1
  elif result2 is not None:
    return result2
  else:
    return None, "Timeout em ambas as APIs"

}