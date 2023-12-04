# Maps

Go calls hash tables as Maps. It has a key and a value.

<a href="https://gobyexample.com/maps" target="_blank">https://gobyexample.com/maps</a>

```go
salary := map[string]int{
  "James": 1000,
  "Alice": 2000,
  "Zack": 3000
}

delete(salary, "Zack")
salary["Benjamin"] = 1500
```

## Declaration

It can be declared like bellow to start an empty map:

`<map name> := make(map[<key type>]<value type>)`
`<map name> := map[<key type>]<value type>{}`

## Traversing the map

```go
salaries := map[string]int{
  "James": 1000,
  "Alice": 2000,
  "Zack": 3000
}

for name, salary := range salaries{
  fmt.Printf("%s's salary is %d.\n", name, salary)
}
```

Output:

> James's salary is 1000.
> Alice's salary is 2000.
> Zack's salary is 3000.

To ignore the key on the for loop add a blank identifier with `_`.

```go
salaries := map[string]int{
  "James": 1000,
  "Alice": 2000,
  "Zack": 3000
}

for _name_, salary := range salaries{
  fmt.Printf("Salary is %d.\n", salary)
}
```

Output:

> Salary is 1000.
> Salary is 2000.
> Salary is 3000.
