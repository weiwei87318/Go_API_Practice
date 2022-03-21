package main

import (
	_ "api_practice/database"
	"api_practice/router"
	//orm "api/database"
)

func main() {
	//defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8080")
}
