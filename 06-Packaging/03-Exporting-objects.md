# Exporting objects

When working with object-oriented languages we can modify the access (protect, public...)

But Go is not an OOL, but we can have something similar to it.

When writing with variable/type/func... in capital letter it will be exported.

```go
type Math struct {
  A int
  b int
}
```

The `Math` struct is being exported. `b` is not being exported.

```go
type math struct {
  A int
  B int
}
```

The `math` struct is not being exported