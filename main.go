package main

import (
	_ "api_practice/database"
	//orm "api/database"
	"api_practice/router"
)

func main() {
	//defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8080")
}
