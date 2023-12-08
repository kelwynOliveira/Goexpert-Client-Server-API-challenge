# Customizing requests

```go
package main

import(
  "io/ioutil"
  "net/http"
  "time"
)

func main(){
  c := http.Client{}
  req, err := http.NewRequest("GET", "https://google.com", nil)
  if err != nil{
    panic(err)
  }
  req.Header.Set("Accept", "application/json")
  ans, err := c.Do(req)
  if err != nil{
    panic(err)
  }
  defer ans.Body.Close()
  body, err := ioutil.ReadAll(ans.Body)
  if err != nil{
    panic(err)
  }
  println(string(body))
}
```
