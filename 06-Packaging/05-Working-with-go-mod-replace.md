# Working with go mod replace

How to work when the package is not on github but on your machine?

```shell
go mod edit -replace <package github name>=<path on your machine>
go mod tidy
```

```shell
go mod edit -replace github.com/devfullcycle/goexpert/7-Packaging=../math
go mod tidy
```