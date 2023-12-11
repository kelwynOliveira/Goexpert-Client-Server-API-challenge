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

  // Selecting more than one product
  products, err := selectAllProducts(db)
  if err != nil{
    panic(err)
  }
	for _, p := range products {
		fmt.Printf("Product: %v, has price %.2f\n", p.Name, p.Price)
	}
}

func selectAllProducts(db *sql.DB) ([]Product, err) {
  rows, err := db.Query("SELECT id, name, price FROM products")
  if err != nil{
    panic(err)
  }
  defer stmt.Close()
  var products []Product
  for rows.Next(){
    var p Product
    err = rows.Scan(&p.ID, &p.Name, &p.Price)
    if err != nil{
      panic(err)
    }

    products = append(products, p)
  }
}
```
