# Structs and methods

We can create methods for `structs`.

<a href="https://gobyexample.com/methods" target="_blank">https://gobyexample.com/methods</a>

## Structure for a struct method

`func (<parameter to call struct> <name of the struct>) <function name>(...){...}`

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
  Address
}

// Method to turn client not active
func (c Client) Inactive(){
  c.Active = false
  fmt.Printf("Client %s is now inactive", c.Name)
}


func main(){
  James := Client{
    Name: "James",
    Age: 29,
    Active: true,
  }
  James.Inactive() // Turn client inactive

  fmt.Printf("Name: %s, Age: %d, Active: %t", James.Name, James.Age, James.Active)
}
```
