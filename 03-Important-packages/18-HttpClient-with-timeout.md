# HttpClient with timeout

```go
package main

import(
  "io/ioutil"
  "net/http"
  "time"
)

func main(){
  c := http.Client{Timeout: time.Second} // time.Duration(1) * time.Second | time.Microsecond
  ans, err := c.Get("https://google.com")
  if err != nil{
    panic(err)
  }
  defer ans.Body.ReadAll(ans.Body)
  body, err := ioutil.ReadAll(ans.Body)
  if err != nil{
    panic(err)
  }
  println(string(body))
}
```
