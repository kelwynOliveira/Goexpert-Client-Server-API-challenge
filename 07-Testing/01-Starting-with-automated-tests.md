# Starting with automated tests


```
go help test
```

> File called tax.go
```go
package tax

func CalculateTax(amount float64) float64 {
    if amount >= 1000{
        return 10.0
    }
    return 5.0
}
```

> File called tax_test.go
The file must end with _test

Make sure the tax_test has the same package as tax
```go
package tax

import "testing"

func TestCalculateTax(t *testing.T) {
    amount := 500
    expected := 5.0
    result := CalculateTax(amount)

    if result != expected {
        t.Errorf("Expected %f but got %f", expected, result)
    }
}
```

Run the test:
```
go test .
```

or to see verbosely
```
go test -v
```