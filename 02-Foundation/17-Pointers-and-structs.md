# Pointers and structs

We use `*` to point to the original value. Otherwise the value on the function will be just a copy.

```go
package main

type Client struct{
  name string
}

func (c *Client) walked(){
  c.name = "James Smith"
  fmt.Printf("The client %v walked", c.name)
}

func main(){
  James := Client{
    name: "James"
  }
  James.walked()

  fmt.Printf("The value of the struct with name is: %v", James.name)
}
```

This below is like a constructor.

```go
type Account struct{
  balance int
}

func NewAccount() *Account{
  return &Account{balance: 0}
}
```
