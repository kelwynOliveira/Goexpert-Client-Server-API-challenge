# Custom data types

Go let us customize or own data types. To do it the struct is: `type <new type name> <base type>`.

`<base type>` to say if it is `int`, `bool`, `string`...

```go
package main

type ID int

var f ID = 1

func main(){
  println(f)
}
```

The new type will support any operation that comes with its base type.

> There are a few reasons you'd generally want to use a new type:
>
> - You need to define methods on the type.
> - You don't want the type to be comparable with literals or variables with the type it's derived from (eg. to reduce user confusion or make sure they don't do something invalid like attempt to compare your special string with some other random string).
> - You just need a place to put documentation, or to group methods that return a specific type (eg. if you have several Dial methods that return a `net.Conn`, you might create a `type Conn net.Conn` and return that instead just for the sake of grouping the functions under the Conn type header in godoc or to provide general documentation for the net.Conn returned by the methods).
> - Because you want people to be able to check if something of a generic type came from your package or not (eg. even if your Conn is just a normal net.Conn, it gives you the option of type switching and checking if it's a yourpackage.Conn as well)
>   You want a function to take an argument from a predefined list of things and you don't want the user to be able to make new values that can be passed in (eg. a list of exported constants of an unexported type).
>
>   <a href="https://stackoverflow.com/questions/40268276/best-practices-to-use-own-type-in-golang-alians-to-build-type" target="_blank">Stack overflow</a>
