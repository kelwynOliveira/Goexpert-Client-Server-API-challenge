# Modules and packages

All go files will start with a package. The package name must be the name of the folder where the file is. Except for the package main.

Suppose the following directory:

- Project/
  -- math/
  --- math.go
  -- main.go

math.go:

```go
package math

func Sum[T int | float64] (a, b T) T{
  return a + b
}
```

main.go:

```go
package main

import (
 "fmt"
 "math"
)

func main() {
  s := math.Sum(10, 20)
  fmt.Printf("The sum is: %v", s)
}
```

After run the result will be: `package math is not in GOROOT (/usr/local/go/src/math)`

It would have worked if the package `math` where in the `GOROOT` because Go will look for every package we import in the `GOROOT`.

## The GOMOD

It is a package manager for Go.

The first step is to create a go mod with `go mod init <module name>` where you want to run the project.

A good practice is to address your github project folder, this way other user can just use it to install your module.

eg.: On cmd

```go
go mod init github.com/kelwynOliveira/go-course
```

It will create a file called go.mod:

```go
module github.com/kelwynOliveira/go-course

go 1.19
```

Now go will not only look on the `GOROOT` but also in the module specified on go.mod:

```go
package main

import (
 "fmt"
 "github.com/kelwynOliveira/go-course/math"
)

func main() {
  s := math.Sum(10, 20)
  fmt.Printf("The sum is: %v", s)
}
```

## Exporting to other packages

The function, variable, struct... name must start Cased:

`func sum(...){...}` will not work outside the package where it was created
`func Sum(...){...}` will work on other files that imports the package.
