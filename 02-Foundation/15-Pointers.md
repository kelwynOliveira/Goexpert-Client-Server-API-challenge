# Pointers

<a href="https://gobyexample.com/pointers" target="_blank">https://gobyexample.com/pointers</a>

```go
a := 10
println(&a)
```

> The variable points to a pointer that has the address in the memory that has the value.

"a" is the memory address where "10" is.
`&a` is the address of the value `a`.

```go
a := 10
var pointer *int = &a
*pointer = 20 // Goes to where pointer is point (address of the value in a)
b := &a // Save the address of where the value of a is
println(*b) // Shows the value of a (deference)
*b = 30 // Goes to the address where b is pointing and change the value
println(a) // Since b is the address of the value a its value is now 30
```

When using `*` it refers to the address.

We use pointers to work to change the value in memory (not local scope).
