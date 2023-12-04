# Type assertion

```go
package main

func main(){
  var hello interface{} = "Hello, World!"

  println(hello)
  println(hello.(string))
  res, ok := hello.(int)
  res2 := hello.(int)
}
```

If the conversion to `int` is not okay `res` will be 0 and `ok` will be `false`.

`res2` will give a `panic` message.
