# Finishing response to HTTP

```go
package main

import (
  "net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main(){
  http.HandleFunc("/", SearchCEPHandler)
  http.ListenAndServe(":8080", nil)
}

func SearchCEPHandler (w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    w.WriteHeader(http.StatusNotFound)
    return
  }
  cepParam := r.URL.Query().Get("cep")
  if cepParam == "" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  cep, err := SearchCEP(cepParam)
  if err != nil{
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  // result, err := json.Marshal(cep)
  // if err != nil{
  //   w.WriteHeader(http.StatusInternalServerError)
  //   return
  // }
  json.NewEncoder(w).Encode(cep)
}

func SearchCEP(cep string) (*ViaCEP, error){
  answer, err := http.Get("https://viacep.com.br/ws/"+ cep + "/jason/")
  if err != nil{
    return nil, err
  }
  defer answer.Body.Close()
  body, err := ioutil.ReadAll(answer.Body)
  if err != nil{
    return nil, err
  }
  var c ViaCEP
  err = json.Unmarshal(body, &c)
  if err != nil{
    return nil, err
  }
  return &c, nil
}
```

Start the server: `curl localhost:8080`
