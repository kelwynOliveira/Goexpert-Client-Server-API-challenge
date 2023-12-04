# Variadic functions

<a href="https://gobyexample.com/variadic-functions" target="_blank">https://gobyexample.com/variadic-functions</a>

## Structure of a variational function

Used when we want to give a number of parameters but don't know how many (eg. parameters to configure a database it could be one or more). We use `...` before the parameter type when declaring a function.

`func <function name>(<parameter name> ...<parameter type>) (<return types>){
  ...
}`

```go
func sum(numbers ...int) int {
  total := 0
  for _, number := range numbers {
    total += number
  }
  return total
}
```

## Calling the function

`<function name>(<parameters values>)`

```go
sum(1, 2, 10, 20, 12)
```
