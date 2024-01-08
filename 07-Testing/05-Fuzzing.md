# Fuzzing

It works as a mutation test

From previous tax.go:
```go
package tax

func CalculateTax(amount float64) float64 {
    if amount < 0{
        return 0
    }
    if amount >= 1000{
        return 10.0
    }
    if amount >= 20000{
        return 20.0
    }
    return 5.0
}
```

From previous tax_test.go:
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

//Testing

//Benchmarking

//Fuzzing
func FuzzCalculateTax(f *testing.F) {
    seed := []float64{-1, -2, -2.5, 500.0, 1000.0, 1501.0}
    for _, amount := range seed {
        f.Add(amount)
    }
    f.Fuzz(func(t *testing.T, amount float64){
        result := CaulculateTax(amount)
        if amount <= 0 && result != 0 {
            t.Error("Received %f but expected 0", result)
        }
        if amount > 20000 && result != 20 {
            t.Error("Received %f but expected 20", result)
        }
    })
}
```

`Seeding` you give some informations to know witch numbers to generate.

## Running
```
go test -fuzz=. -run=^#
```

It creates a folder called "testdata" with infos on were it is failing.