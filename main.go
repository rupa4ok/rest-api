package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	Config2 "main/src/Config"
	Models2 "main/src/Models"
	Routers2 "main/src/Routers"
)

var err error

func main() {

	Config2.DB, err = gorm.Open("mysql", "user:password@tcp(rest-go-mysql:3306)/db?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config2.DB.Close()
	Config2.DB.AutoMigrate(&Models2.Book{})

	r := Routers2.SetupRouter()
	// running
	r.Run()
}
