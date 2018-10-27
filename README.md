# skeleton-go-gin-gorm

# Api go simple
It is a just simple skeleton Go using:
1. **Gin Framework**
2. **Gorm** 
3. **Mysql**

## Installation & Run
```bash
# Download this project
$ go get github.com/dedidot/skeleton-go-gin-gorm

# Download Gin Framework
$ go get github.com/gin-gonic/gin

# Download GORM
$ go get github.com/jinzhu/gorm

# Mysql
$ go get github.com/go-sql-driver/mysql

```

Setting DB in config/database.go
```go
Config.DB, err = gorm.Open("mysql", "DBUSERNAME:DBPASSWORD@tcp(127.0.0.1:3306)/DBNAME?charset=utf8&parseTime=True&loc=Local")
```

## Structure
```
├── app
│   ├──handler   
│     └── simple.go // controller
|	  ├── model.go 
|     └── model.go // model, get data from database 
|     └── scheme.go // tabel list
|   ├── helper
│     └── Response.go // response function
|   └── routers.go // Routers
├── config
│   └── Database.go // Global DB
└── main.go
```

you can try make crud with golang, look my other repo: https://github.com/dedidot/gorm-gin
