# For loop

Go only have the `for` loop.

<a href="https://gobyexample.com/for" target="_blank">https://gobyexample.com/for</a>

## For structure

```go
for i := 0; i < 10; i++{
  println(i)
}
```

```go
i := 0
for i < 10{
  println(i)
  i++
}
```

```go
numbers := []string{"one", "two", "three"}

for k, v := range numbers {
  fmt.Printf("Key %v, value %s.\n", k, v)
}
```

Infinity loop

```go
for {
  ...
}
```
