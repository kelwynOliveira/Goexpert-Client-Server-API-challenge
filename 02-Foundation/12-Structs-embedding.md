# Structs embedding

<a href="https://gobyexample.com/struct-embedding" target="_blank">https://gobyexample.com/struct-embedding</a>

```go
package main

import "fmt"

type Address struct {
  Street  string
  Number  int
  City    string
  State   string
}

type Client struct {
  Name    string
  Age     int
  Active  bool
  Address // Address is embedded into the Client struct
}

func main(){
  James := Client{
    Name: "James",
    Age: 29,
    Active: true,
  }
  James.City = "Any place"
  James.Address.State = "A bigger place"

  fmt.Printf("Name: %s, Age: %d, Active: %t", James.Name, James.Age, James.Active)
}
```

We can also create a variable using the `struct` as a type:

```go
package main

type Address struct {
  Street  string
  Number  int
  City    string
  State   string
}

type Client struct {
  Name    string
  Age     int
  Active  bool
  Place   Address // Place is a property of type Address
}

func main(){
  ...
}
```
