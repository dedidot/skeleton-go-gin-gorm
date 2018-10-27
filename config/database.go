package config

import(
	"fmt"
	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
	err error
)

func init() {
	DB, err = gorm.Open("mysql", "USERNAME:PASSWORD@tcp(127.0.0.1:3306)/DATABASENAME?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer DB.Close()
}