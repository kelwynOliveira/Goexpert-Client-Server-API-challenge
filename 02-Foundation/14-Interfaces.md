# Interfaces

Go interfaces only accepts methods.

<a href="https://gobyexample.com/interfaces" target="_blank">https://gobyexample.com/interfaces</a>

```go
package main

import "fmt"

type Address struct {
  Street  string
  Number  int
  City    string
  State   string
}

//
type Person interface {
  Inactive()
}

type Client struct {
  Name    string
  Age     int
  Active  bool
  Address
}

func (c Client) Inactive(){
  c.Active = false
  fmt.Printf("Client %s is now inactive", c.Name)
}

//
func Disable(person Person){
  person.Inactive()
}

func main(){
  James := Client{
    Name: "James",
    Age: 29,
    Active: true,
  }
  Disable(James) //
}
```
