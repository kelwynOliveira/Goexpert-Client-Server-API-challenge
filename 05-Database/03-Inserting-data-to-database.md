# Inserting data to database

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

  product := NewProduct("Notebook", 18999.9)
  err = insertProduct(db, product)
  if err != nil{
    panic(err)
  }
}

func insertProduct(db *sql.DB, product *Product) error {
  stmt, err := db.Prepare("insert into products(id, namee, price) values(?, ?, ?)")
  if err != nil{
    panic(err)
  }
  defer stmt.Close()
  _, err = stmt.Exec(product.ID, product.Name, product.Price)
  if err != nil {
    panic(err)
  }
  return nil
}
```

> mysql driver `github.com/go-sql-driver/mysql`

`_` means blank identifier. To maintain the package.

`root:root@tcp(localhost:3306)/goexpert` = `user:password(localhost:ports)/project`
