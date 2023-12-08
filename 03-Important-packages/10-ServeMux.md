# ServeMux

```go
package main

import (
  "net/http"
)

func main(){
  mux := http.NewSrveMux()
  // mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
  //   w.Write([]byte("hello world!"))
  // })
  mux.HandleFunc("/", HomeHandler)
  mux.Handle("/blog", blog{title: "My Blog"})
  http.ListenAndServe(":8080", nil)

  mux2 := http.NewServeMux()
  mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("hello world!"))
  })
  http.ListenAndServe(":8081", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello, World!"))
}

type blog struct{
  title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Blog"))
}
```
