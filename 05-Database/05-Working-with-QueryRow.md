# Working with QueryRow

```go
package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/google/uuid"
)

type Product struct{
  ID string
  Name string
  Price float64
}

func NewProduct(name string, price float64) *Product{
  return &Product{
    ID: uuid.New().String(),
    Name: name,
    Price: price,
  }
}

func main(){
  db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
  if err != nil{
    panic(err)
  }
  defer db.Close()

  // ...

  // Selecting
  p, err := selectProduct(db, product.ID)
  if err != nil{
    panic(err)
  }
  fmt.Printf("Product: %v, has price %.2f", p.Name, p.Price)
}


func selectProduct(db *sql.DB, id string) (*Product, err){
  stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
  if err != nil{
    panic(err)
  }
  defer stmt.Close()
  var p Product
  err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
  if err != nil{
    panic(err)
  }
  return &p, nil
}
```
