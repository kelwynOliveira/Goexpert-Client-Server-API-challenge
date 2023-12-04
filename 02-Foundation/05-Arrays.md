# Arrays

Array works as a variable, it has a fixed size and can be traversed.

<a href="https://gobyexample.com/arrays" target="_blank">https://gobyexample.com/arrays</a>

## Declaration

To declare an Array variable we do `var <array name> [<array size>]<array type>`

The example bellow is called "myArray", its size is 3 (have 3 items) and the values are of the `int` type:

```go
package main

import "fmt"

var myArray [3]int
myArray[0] = 1
myArray[1] = 2
myArray[1] = 3

func main(){
  ...
}
```

If you input more than what is the size specified Go will say it is out of bound. It means that a array has a fixed size and cannot input more items than what was specified.

## Accessing the array

It can be accessed by its position.

> eg.: `myArray[2]`. It will access the item in position 2 of the array "myArray".

## Traversing the array

An array can be traversed. One way to do it is using a for loop.

```go
package main

import "fmt"

var myArray [3]int
myArray[0] = 10
myArray[1] = 20
myArray[1] = 30

func main(){
  for i, v := range myArray{
    fmt.Printf("The index value is %d and the value is %d.\n", i, v)
  }
}
```

Output:

> The index value is 0 and the value is 10
> The index value is 1 and the value is 20
> The index value is 2 and the value is 30

`range` will traverse the array.
