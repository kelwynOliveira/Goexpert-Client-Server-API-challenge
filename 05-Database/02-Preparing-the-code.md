# Preparing the code

Normaly we try to use everything that comes with the Go language instead of using libraries.
Although Go has the GORM(Go ORM) we do SQL.

File docker-compose.yaml:

```yaml
version: '3'
services:
  mysql:
  mysql:
    image: mysql:5.7
    container_name: mysql
    restart: always
    platform: linux/amd64
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: goexpert
      MYSQL_PASSWORD: root
    ports:
      - 3306:3305
```

`docker-compose up -d`
`docker-compose exec mysql bash`
`mysql -uroot -p goexpert`
`root`

> SQL table:
> `CREATE TABLE products (id varchar(255), name varchar(80), price decimal(10,2), primary key (id));`

```go
package main

import (
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

}
```

`go mod init <github repo>`
`go mod tidy`
