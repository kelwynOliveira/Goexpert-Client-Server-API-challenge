# HTTP calls

```go
package main

import (
  "io"
  "net/http"
)

func main(){
  request, err := http.Get("https://www.google.com/")
  if err != nil {
    panic(err)
  }
  result, err := io.ReadAll(request.Body)
  if err != nil {
    panic(err)
  }

// Print as string
  println(string(request))

// close teh request
  request.Body.Close()
}
```
