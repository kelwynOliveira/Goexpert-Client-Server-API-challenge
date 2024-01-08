# Starting with Testify

From previous tax.go:
```go
package tax

func CalculateTax(amount float64) (float64, error) {
    if amount <= 0{
        return 0.0, "Must be greater than 0"
    }
    if amount >= 1000{
        return 10.0, nil
    }
    if amount >= 20000{
        return 20.0, nil
    }
    return 5.0, nil
}
```

File `tax_test.go`
From previous tax.go:
```go
package tax

import (
    "testing"
    "github.com/stretchr/testfy/assert"
)

func TestCalculateTax(t *testing.T) {
    tax := CalculateTax(1000.0)
    assert.Nil(t, err)
    assert.Equal(t, 10.0, tax)

    tax := CalculateTax(0.0)
    assert.Error(t, err, "Must be greater than 0")
    assert.Equal(t, 0.0, tax)
}
```

There is a lot of assert types