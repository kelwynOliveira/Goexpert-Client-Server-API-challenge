# Working with Bentchmarking

`T` is for testing (unit testing)
`B` is for benchmarking (how many will be used on the benchmark)

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
func BenchmarkCalculateTax(b *testing.B) {
    for i := 0; i < b.N; i++ {
        CalculateTax(500.0)
    }
}
```

## run:
```
go test -bench=.
```

CPU quantity used
> BenchmarkCalculateTax-20

Times runnig the test
> 100000000

Operations runned per nano seconds
> 0.3710 ns/op

## run only the benchmark
```
go test -bench=. -run=^#
```
give a regex to run the test that starts with `#`

Can be used to compare algorithms.


## run the benchmark for a specific time
```
go test -bench=. -run=^# -count=10
```

## run with a specific time
```
go test -bench=. -run=^# -count=10 -benchtime=3s
```

## run to see related to memory allocation
```
go test -bench=. -run=^# -benchmem
```