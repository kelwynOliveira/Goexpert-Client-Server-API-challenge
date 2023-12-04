# Generics

<a href="https://gobyexample.com/generics" target="_blank">https://gobyexample.com/generics</a>

> Starting with version 1.18, Go has added support for generics, also known as type parameters.

We can have generics in our functions and in the structs methods.

Consider the following:

```go
package main

func SumInteger(m map[string]int) int {
  var sum int
  for _, v := range m {
    sum += v
  }
  return sum
}

func SumFloat(m map[string]float64) float64 {
  var sum float64
  for _, v := range m {
    sum += v
  }
  return sum
}

func main(){
  salariesIntegers := map[string]int{
    "James": 1000,
    "Alice": 2000,
    "Zack": 3000
  }
  salariesFloats := map[string]int{
    "James": 1000.0,
    "Alice": 2000.0,
    "Zack": 3000.0
  }
  println(SumInteger(salariesIntegers))
  println(SumFloat(salariesFloats))
}
```

It is not a good case to create a second function with the same activities just because it has different types.

With generics the above will be:

```go
package main

func Sum[T int | float64] (m map[string]T) T {
  var sum T
  for _, v := range m {
    sum += v
  }
  return sum
}

func main(){
  salariesIntegers := map[string]int{
    "James": 1000,
    "Alice": 2000,
    "Zack": 3000
  }
  salariesFloats := map[string]int{
    "James": 1000.0,
    "Alice": 2000.0,
    "Zack": 3000.0
  }
  println(Sum(salariesIntegers))
  println(Sum(salariesFloats))
}
```

## Constraints

The constraints package: <a href="https://pkg.go.dev/golang.org/x/exp/constraints" target="_blank">https://pkg.go.dev/golang.org/x/exp/constraints</a>

```go
package main

// Constraint
type Number interface{
  int | float64
}

func Sum[T Number] (m map[string]T) T {
  var sum T
  for _, v := range m {
    sum += v
  }
  return sum
}

func main(){
  salariesIntegers := map[string]int{
    "James": 1000,
    "Alice": 2000,
    "Zack": 3000
  }
  salariesFloats := map[string]int{
    "James": 1000.0,
    "Alice": 2000.0,
    "Zack": 3000.0
  }
  println(Sum(salariesIntegers))
  println(Sum(salariesFloats))
}
```

## Specifying a type

```go
package main

// My type
type MyNumber int

type Number interface{
  int | float64
}

func Sum[T Number] (m map[string]T) T {
  var sum T
  for _, v := range m {
    sum += v
  }
  return sum
}

func main(){
  salariesIntegers := map[string]int{
    "James": 1000,
    "Alice": 2000,
    "Zack": 3000
  }
  salariesFloats := map[string]int{
    "James": 1000.0,
    "Alice": 2000.0,
    "Zack": 3000.0
  }
  salariesMyNumber := map[string]MyNumber{
    "James": 1000,
    "Alice": 2000,
    "Zack": 3000
  }
  println(Sum(salariesIntegers))
  println(Sum(salariesFloats))
  println(Sum(salariesMyNumber))
}
```

Despite `MyNumber` been the same type as `int` for been strongly typed Go will not let you do this. It sees `int` and `MyNumber` as different types.

To make it work we just need to add `~` on the constraint.

```go
type Number interface{
  ~int | ~float64
}
```

## The `comparable` constraint

The `comparable` constraint will let you use only if the values can be compared (same type).

```go
package main

func Compare[T comparable](a, b T)bool{
  if a == b{
    return true
  }
  return false
}

func main(){
  println(Compare(10, 10))
}
```

The following (if statement) will not be possible because `compare` cannot be certain if the values are comparable in higher or lower than:

```go
package main

func Compare[T comparable](a, b T)bool{
  if a > b{
    return true
  }
  return false
}

func main(){
  ...
}
```

> The `comparable` constraint, means that we can compare values of this type with the `==` and `!=` operators.

## The any constraint

`any` is an alias for interface{} (empty interface).

```go
package main

func Compare[T any](a, b T)bool{
  ...
}

func main(){
  ...
}
```
