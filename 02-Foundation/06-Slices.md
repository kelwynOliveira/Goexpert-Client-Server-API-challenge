# Slices

Slices are arrays and have three main characteristics:

1. They have a pointer to point to the array.
1. Has a size: to know until where it can go.
1. Has a capacity: to know how much information it can receive.

<a href="https://gobyexample.com/slices" target="_blank">https://gobyexample.com/slices</a>

## Attribution

```go
s := []int{2, 4, 6, 8, 10}
```

## Slicing

```go
s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

s[:4]
```

`s[:4]` will slice the array (drop) the values that comes after position 4. Its capacity will be the same as the initial capacity (10) but it will have only 4 values (10, 20, 30, 40).

```go
s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

s[4:]
```

`s[4:]` will slice the array (drop) the values that came before position 4. Its capacity will be different, the capacity will be (6) and it will have 6 values (50, 60, 70, 80, 90, 100).

## Adding items to the slice

Slice works with arrays. So when it needs to use a bigger array it just points to a bigger array.

```go
s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

s = append(s, 110)
```

The capacity will be (18) and it will have 11 values (10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110).
