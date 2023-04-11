package main

import (
	"be-ifid/config"
	"be-ifid/database"
	"fmt"
)

func init() {
	config.Init()
	database.Init()
}

func main() {
	fmt.Println("Hello World")
}
