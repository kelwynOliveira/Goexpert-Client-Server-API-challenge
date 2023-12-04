# Starting with structs

Since Go is not an Object-Oriented language "Go has its 'Go way to program'".

The `Structs` is the closer we have in Go to work like an Object-oriented language. A `struct` is not a `class`.

<a href="https://gobyexample.com/structs" target="_blank">https://gobyexample.com/structs</a>

## Structure of a struct

```go
package main

import "fmt"

type Client struct {
  Name string
  Age int
  Active bool
}

func main(){
  James := Client{
    Name: "James",
    Age: 29,
    Active: true,
  }

  fmt.Printf("Name: %s, Age: %d, Active: %t", James.Name, James.Age, James.Active)
}
```

> We can use a `struct` to compose another struct, but we cannot work with heritage.
