package main

import (
	"fmt"
	"github.com/Jchann24/go-gorm-rest-api/Configs"
	"github.com/Jchann24/go-gorm-rest-api/Models"
	"github.com/Jchann24/go-gorm-rest-api/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
