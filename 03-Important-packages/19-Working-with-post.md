# Working with post

```go
package main

import(
  "io/ioutil"
  "net/http"
  "time"
)

func main(){
  c := http.Client{}
  jsonVar := byte.NewBuffer([]byte(`{"name": "my name"}`)) // Need to buffer the data before posting it
  ans, err := c.Post("https://google.com", "application/json", jsonVar)
  if err != nil{
    panic(err)
  }
  defer ans.Body.ReadAll(ans.Body)
  io.CopyBuffer(os.Stdout, ans.Body, nil)
  println(string(ans))
}
```
