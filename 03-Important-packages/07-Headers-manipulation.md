# Headers manipulation

```go
package main

import (
  "net/http"
)

func main(){
  http.HandleFunc("/", SearchCEPHandler)
  http.ListenAndServe(":8080", nil)
}

func SearchCEPHandler (w http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    w.WriteHeader(http.StatusNotFound)
    return
  }

  // ask cep in query (?cep=12345678910)
  cepParam := r.URL.Query().Get("cep")
  if cepParam == "" {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello, World!"))
}
```

Start the server: `curl localhost:8080`
