# Attribution and declaration

Go is a strongly typed language. It means that if a variable is declared as `bool` it cannot have its type changed to `int` or other variable type.

If you have a **local variable** and it is not been used, the compiler will prompt a message saying that a variable is not been used and will not compile your code until you use the variable or remove it.

## Constant

It cannot have its value changed.

```go
const a = "Hello, World!"
```

## Variables declaration

It can have its value changed.

```go
var b bool
var c int
var d string
var e float64
```

```go
var (
  b bool
  c int
  d string
  e float64
)
```

If we do not do a value attribution it will be inferred a value.

| type      | Inferred value   |
| --------- | ---------------- |
| `bool`    | `false`          |
| `int`     | `0`              |
| `string`  | ` `              |
| `float64` | `+0.000000e+000` |
| `float32` | `+0.`            |
| `complex` |                  |
| `byte`    |                  |
| `rune`    |                  |

## Variables attribution

We can declare the type and attribute a value.

```go
var b bool = true
var c int = 10
var d string = "Hello, World!"
var e float64 = 10.1
```

```go
var (
  b bool = true
  c int = 10
  d string = "Hello, World"
  e float64 = 10.1
)
```

# Scope of the variable

If the variable declaration was outside the methods it can be accessed for any method (global variable):

```go
package main

var a = "Hello, World!"

func main(){
  println(a)
}

func sayHello(){
  println(a)
}
```

But if the variable have its declaration inside a method (local variable) it cannot be accessed for other methods:

```go
package main

func main(){
  var a = "Hello, World!"
  println(a)
}

func sayHello(){
  var a = "Hello, World!"
  println(a)
}
```

## Shorthand declaration

If we want to declare a variable and make its value attribution in Go we have a shorthand for it.

Instead of doing `var a string = "Hello, World!"` we can do `a := "Hello, World!"`.

We use `:=` to say it is a new variable with a value. Go will infer the type according to the value of the variable.
