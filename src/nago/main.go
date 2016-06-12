package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("toml")
	viper.AddConfigPath("/Users/lisces/Code/nago/src/nago")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(viper.Get("database.host"))
	fmt.Println(viper.Get("database.port"))
	fmt.Println(viper.Get("database.name"))
	fmt.Println(viper.Get("database.user"))
	fmt.Println(viper.Get("database.pass"))
	fmt.Println("test")
}
