# Understanding the first line

All go files will start with a package. The package name must be the name of the folder where the file is. Except for the package main.

> main.go file:

```go
package main

func main(){
  println("Hello, World!")
}
```

Suppose you have a directory like this:
\- Project/
\-- sum/
\--- s.go
\-- main.go

Inside the sum/ folder we have a file called `s.go`. Since it is inside the "sum folder" it must start with `package sum`, the name of the folder.

> Inside the s.go file:

```go
package sum

...
```

If there is more than one file inside the folder, these files must have the package with folder's name.

Suppose you have a directory like this:
\- Project/
\-- operations/
\--- sum.go
\--- subtraction.go
\-- main.go

> Inside the `sum.go` and subtraction.go file we must have:

```go
package operations

...
```

Except for the main package. e.g.: Inside `sum.go` or `subtraction.go` we can have:

```go
package main

...
```

When inside the same folder the files with the same package will have access to each other.

Suppose you have a directory like this:
\- Project/
\-- main.go
\-- sum.go

Both `main.go` and `sum.go` uses the `package main`. In this case the `main.go` file will have access to everything inside `sum.go` (variables, functions...) and `sum.go` will have access to everything inside `main.go` because they are at the same directory level and have the same package.
