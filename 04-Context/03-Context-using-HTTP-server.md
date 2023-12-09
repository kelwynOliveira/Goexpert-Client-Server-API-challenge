# Context using HTTP server

```go
package main

import (
  "net/http"
)

func main(){
  http.HandlerFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
  ctx := r.Context()
  log.Println("Requesting...")
  defer log.Println("Requeste finished.")
  select{
    case <- time.After(5 * time.Second):
    // Prints on stdout
      log.Println("Request succeeded.")
      // Prints on browser
      w.Write([]byte("Request cancelled."))
    case <- ctx.Done():
    // Prints on stdout
      log.Println("Request cancelled by client")
  }
}
```
