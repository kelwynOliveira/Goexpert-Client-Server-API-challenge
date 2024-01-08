# Verifying code coverage

```
go test -coverprofile=coverage.out
```
It shows the coverage of the test.
It generates a file called coverage.out

We use the go tool with:
```
go tool cover -html=coverage.out
```
