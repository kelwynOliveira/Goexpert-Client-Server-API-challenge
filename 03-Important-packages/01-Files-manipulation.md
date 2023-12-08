# Files manipulation

The OS package: <a href="https://pkg.go.dev/os" target="_blank">https://pkg.go.dev/os</a>

```go
package main

import (
  "fmt"
  "os"
)

func main(){
  f, err := os.Create("file.txt")
  if err != nil {
    panic(err)
  }

// Writing
  // size, err := f.WriteString("Hello, World!")
  size, err := f.Write([]byte("Writing"))
  if err != nil {
    panic(err)
  }
  fmt.Printf("File created with success! Size: %d bytes.\n", size)
  f.close

// Opening
  file, err := os.Open("file.txt")
  if err != nil {
      panic(err)
  }

// Reading
  file, err := os.ReadFile("file.txt")
  if err != nil {
      panic(err)
  }
  fmt.Println(string(file))
}
```

## Reading file bigger than the memory

> Consider we have only 100MB of memory but we want to read a file larger than this. Instead of load all the file to the memory Go let us read only parts of the file.

```go
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main(){
// Reading
  file, err := os.Open("file.txt")
  if err != nil{
    panic(err)
  }
  reader := bufio.NewReader(file)
  buffer := make([]byte, 10)
  for{
    n, err := reader.Read(buffer)
    if err != nil {
      break
    }
    fmt.Println(string(buffer[:n]))
  }
}
```

## Removing a file

```go
package main

import (
  "os"
)

func main(){
  err := os.Remove("file.txt")
  if err != nil{
    panic(err)
  }
}
```

| Command      | Function         |
| ------------ | ---------------- |
| `Create`     | to create a file |
| `Open`       | to open a file   |
| `ReadFile`   | to read a file   |
| `WriteStrig` | writes a string  |
| `Write`      | writes data      |
