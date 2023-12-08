# Working with HTTP using contexts

`context` is a package that let us transfer its informations to various calls in our system and we have the option to cancel/stop activities according to the contexts.

`context` normally is a given time.

```go
package main

import()

func main(){
  ctx := context.Background()
  ctx, cancel := context.WithTimeout(ctx, time.Second)
  defer cancel()

  req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
  if err != nil{
    panic(err)
  }
  ans, err := http.DefaultClient.Do(req)
  if err != nil{
    panic(err)
  }
  defer ans.Body.Close()
  body, err := ioutil.ReadAll(ans.Body)
  if err != nil{
    panic(err)
  }
  println(body)
}
```

> ctx, cancel := context.WithCancel(ctx)
