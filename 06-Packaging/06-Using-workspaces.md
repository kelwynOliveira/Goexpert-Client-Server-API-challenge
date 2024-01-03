# Using workspace

Inside the root folder
```shell
go work init <folder1> <folder2>
```

It will create a go.work file with the relatives packages on your machine.

How to use workspace and repo?
1. publish it on github
1. run `go mod tidy -e` will ignore the packages that were not found.