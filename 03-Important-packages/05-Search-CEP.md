# Search CEP

```go
package main

import (
  "io"
  "os"
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
  for _, cep := range os.Args[1:]{

// Resquesting json
    request, err := http.Get("https://www.viacep.com.br/ws/"+ cep +"/json/")
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error accessing %v.\n", err)
    }
    defer request.Body.Close()

// Reading json
    result, err := io.ReadAll(request.Body)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error reading result: %v.\n", err)
    }

// Creating the struct
    var data ViaCEP
    err = json.Unmarshal(result, &data)
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error parsing result: %v.\n", err)
    }

// Printing the struct
    fmt.Println(data)

// Creating file to save data
    file, err := os.Create("city.txt")
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error creating the file: %v.\n", err)
    }
    defer file.Close()

// Wrinting in flie
    _, err := file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.CEP, data.Localidade, data.Uf))

    fmt.Println("Success creating the file!")
    fmt.Println("City: %s", data.Localidade)
  }
}
```
