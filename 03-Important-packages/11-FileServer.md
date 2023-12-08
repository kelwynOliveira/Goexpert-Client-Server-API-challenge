# FileServer

```go
package main

import (
  "log"
  "net/http"
)

func main(){
  fileServe := http.FileServer(http.Dir("./public"))
  mux := http.NewServeMux()
  mux.Handle("/", fileServe)

  log.Fatal(http.ListenAndServe(":8080", mux))
}
```
