package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type User struct {
	id      int
	name    string
	pass    string
	created string
	updated string
}

func main() {
	viper.SetConfigType("toml")
	viper.AddConfigPath("/Users/lisces/Code/nago/src/nago")
	viper.AddConfigPath("/home/vagrant/code/nago/src/nago")
	viper.SetConfigName("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// var dsn = "root:sonata@tcp(localhost:3306)/st?charset=utf8mb4"
	var dsn = viper.Get("database.user").(string)
	dsn += ":"
	dsn += viper.Get("database.pass").(string)
	dsn += "@tcp("
	dsn += viper.Get("database.host").(string)
	dsn += ":"
	dsn += viper.Get("database.port").(string)
	dsn += ")/"
	dsn += viper.Get("database.name").(string)
	dsn += "?charset="
	dsn += viper.Get("database.code").(string)

	db, e := sql.Open("mysql", dsn)
	if e != nil {
		panic(e.Error())
	}

	defer db.Close()

	rows, e := db.Query("SELECT * FROM `users` WHERE 1")
	if e != nil {
		panic(e.Error())
	}

	if rows == nil {
		fmt.Println("rows is nil")
		return
	}

	for rows.Next() {
		user := new(User)
		e := rows.Scan(&user.id, &user.name, &user.pass, &user.created, &user.updated)
		if e != nil {
			fmt.Println("row error")
			return
		}
		fmt.Println(user)
	}
}
