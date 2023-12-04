# Empty interfaces

An empty interfaces accepts any type of data. It is not that good to use because it goes against the highly typed characteristic of Go.

```go
package main

func main(){
  var x interface{} = 10
  var y interface{} = "Hello, World!"

  showType(x)
  showType(y)
}

func showType(t interface{}){
  fmt.Printf("The type of the variable is %T and the value is %v.\n", t, t)
}
```
