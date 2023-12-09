# Context on client side

```go
package main

import (
  "net/http"
  "context"
  "time"
  "os"
)

func main(){
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080", nil)
  if err != nil {
    panic(err)
  }
  defer res.Body.Close()
  io.Copy(os.Stdout, res.Body)
}
```
