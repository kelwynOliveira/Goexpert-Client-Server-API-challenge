# changing and removing registers

```go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main(){
  dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//-- Searching first product
	// var p Product
	// db.First(&p, 1)

  	//-- Making changes
	// p.Name = "New Mouse"
	// db.Save(&p)

  	//-- Searching first product
	// var p2 Product
	// db.First(&p2, 1)

  	//-- Printing product with changes
	// fmt.Println(p2.Name)

  	//-- Removing
	// db.Delete(&p2)
}
```
