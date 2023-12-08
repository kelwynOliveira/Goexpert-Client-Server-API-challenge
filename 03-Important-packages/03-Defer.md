# Defer

Command `defer` will postpone an action

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
  // Since I know it will have to be close at the end
  defer request.Body.Close()

  result, err := io.ReadAll(request.Body)
  if err != nil {
    panic(err)
  }

  println(string(request))
}
```
