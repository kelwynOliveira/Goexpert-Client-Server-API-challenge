# When use pointers

> When we want to change the original value, not just make a copy of the value.

In a function the variable is just a copy of the value passed to the function.

```go
package main

func sum(a, b int) int{
  a = 50
  return a + b
}

func main(){
  var1 := 10
  var2 := 20
  sum(var1, var2)
  println(var1)
}
```

- `sum(var1, var2)`: The returned value is 70. Inside `sum()` a is a copy of the value of `var1` and has its value changed to `50`.
- `println(var1)`: Prints `10`.

To change the value use pointer, to pointer to where the value is in the memory.

```go
package main

func sum(a, b *int) int{
  *a = 50
  return *a + *b
}

func main(){
  var1 := 10
  var2 := 20
  sum(&var1, &var2)
  println(var1)
}
```

- `sum(&var1, &var2)`: Sents the address where the values of `var1` and `var2`is. It returns the value 70. Inside `sum()` `a` is the address where `var1` is on the memory.
- `println(var1)`: Prints `50` since the `func sum` goes to the memory address where `var1` value is been keep.
