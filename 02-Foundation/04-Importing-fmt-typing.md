# Importing `fmt` and typing

To import a package we do `import "<package name>"`

For the `fmt` package that comes from `package main` we do:

```go
package main

import "fmt"

var a = "Hello, World!"

func main(){
  fmt.Printf("The type of a is %T", a)
}
```

The `fmt` comes from "format" as its name implies, this package let us work, it has some resources, with formatting values.

| Item | Prints         |
| ---- | -------------- |
| `%T` | Variable type  |
| `%v` | Variable value |
| `%d` | Digit          |

# Importing more than one package

To import more than one package we can do:

```go
import (
  "fmt"
  "net/http"
)

...
```

> If you are not using a package that was declared, Go will ask to remove it since it is not been used.
