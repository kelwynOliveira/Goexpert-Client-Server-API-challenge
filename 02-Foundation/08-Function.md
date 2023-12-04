# Functions

<a href="https://gobyexample.com/functions" target="_blank">https://gobyexample.com/functions</a>

<a href="https://gobyexample.com/multiple-return-values" target="_blank">https://gobyexample.com/multiple-return-values</a>

## Structure of a function

`func <function name>(<parameters name> <parameters types>) (<return types>){
  ...
}`

```go
func sum(a int, b int) int {
  return a + b
}
```

## Calling a function

`<function name>(<parameters values>)`

```go
sum(1, 2)
```

## Return values

Functions in Go can return more than one value. Its largely used when we think about errors.

```go
func sum(a, b int) (int, bool){
  if a + b >= 50{
    return a + b, true
  }
  return a + b, false
}
```

## Receiving the returned values from a function

```go
package main

import (
  "errors" // Package to create error messages
  "fmt"
)

func main(){
  value, err := sum (50, 10) // "value" receives the 1# returned value and "err" receives the 2#
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(valor)
}

// Function returning a error message
func sum(a, b int) (int, error){
  if a + b >= 50 {
    return 0, errors.New("The sum is greater than 50")
  }
  return a + b, nil
}
```
