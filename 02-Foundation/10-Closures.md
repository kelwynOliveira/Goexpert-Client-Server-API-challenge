# Closures

Also called anonym function. Basically it is a function inside a function. Don't need to be called (specially because they don't have a name [anonym]) it will start by themselves.

<a href="https://gobyexample.com/closures" target="_blank">https://gobyexample.com/closures</a>

## How to do it

```go
package main

import (
  "fmt"
)

func main(){
  total := func() int {
    return sum(1, 2, 10, 20, 12) * 2
  }()

  fmt.Println(total)
}

// Function returning a error message
func sum(a, b int) int{
  if a + b >= 50 {
    return 0
  }
  return a + b
}
```
