# Removing registers

# Searching multiple registers

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

  // Removing
  err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func deleteProduct(db *sql.DB, id string) error {
  stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
  if err != nil{
    panic(err)
  }
  defer stmt.Close()
  _, err = stmt.Exec(id)
  if err != nil {
    panic(err)
  }
  return nil
}


```
