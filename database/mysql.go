package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//model.RedisInit()

var Eloquent *gorm.DB

func init() {
	var err error
	Eloquent, err = gorm.Open("mysql", "root:123456@tcp(mysql:3306)/member?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	err = Eloquent.DB().Ping()
	if err != nil {
		fmt.Printf(err.Error())
	}
	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}
