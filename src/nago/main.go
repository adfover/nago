package main

import (
	"config"
	"fmt"
)

func main() {
	config.GetConfig()
	config.SetConfig()
	fmt.Println("test")
}
