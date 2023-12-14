# Working with JSON

## Struct to JSON

```go
package main

import "enconding/json"

type Account struct{
  Number int
  Balance int
}

func main(){
  account := Account{Number: 1, Balance: 100}
  result, err := json.Marshal(account)
  if err != nil{
    panic(err)
  }
  println(string(result))
}
```

- `Marshal` saves the value

## Sending the result

```go
package main

import (
  "enconding/json"
  "os"
)

type Account struct{
  Number int
  Balance int
}

func main(){
  account := Account{Number: 1, Balance: 100}

// encoder := json.NewEncoder(os.Stdout)
  // encoder.Encode(account)
  err = json.NewEncoder(os.Stdout).Encode(account)
  if err != nil{
    panic(err)
  }
}
```

- `Encode` takes the value, serialize it and deliver/returns it to someone.
- `Decode`

## JSON to struct

```go
package main

import (
  "enconding/json"
)

type Account struct{
  Number int
  Balance int
}


func main(){
  jsonFile := []byte(`{"Number": 2, "Balance": 200}`)
  var accountX Account

  err = json.Unmarshal(jsonFile, &accountX)
  if err != nil{
    panic(err)
  }
  println(accountX.Balance)
}
```

## JSON to struct

When the keys of the JSON and struct does not match transform it will not work. Go has a way to create `tags` to match.

```go
package main

import (
  "enconding/json"
)

type Account struct{
  Number int `json:"n"`
  Balance int `json:"b"`
}

func main(){
  jsonFile := []byte(`{"n": 2, "b": 200}`)
  var accountX Account
  err = json.Unmarshal(jsonFile, &accountX)
  if err != nil{
    panic(err)
  }
  println(accountX.Balance)
}
```
