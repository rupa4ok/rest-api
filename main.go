package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	Config2 "main/src/Config"
	Models2 "main/src/Models"
	Routers2 "main/src/Routers"
	"os"
	"runtime"
)

var err error

func main() {
	fmt.Println("MaxParallelism: ", MaxParallelism())

	dbConnection := os.Getenv("DB_CONNECTION")

	Config2.DB, err = gorm.Open("mysql", dbConnection)

	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer Config2.DB.Close()
	Config2.DB.AutoMigrate(&Models2.Book{})

	r := Routers2.SetupRouter()
	// running
	r.Run()
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}
