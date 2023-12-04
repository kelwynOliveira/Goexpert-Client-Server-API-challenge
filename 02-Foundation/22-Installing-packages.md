# Installing packages

To install we use `go get <module name to get>`, but first we need to create a module with `go mod init <module name>`.

It generates a `go.sum` file.

`go mod tidy` verifies the use of packages and remove it if its not been used or add if it was not downloaded. `go mod tidy` optimizes your `go mod`.
