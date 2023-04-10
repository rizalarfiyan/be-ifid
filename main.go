package main

import (
	"be-ifid/config"
	"fmt"
)

func init() {
	config.Init()
}

func main() {
	fmt.Println("Hello World!")
}
