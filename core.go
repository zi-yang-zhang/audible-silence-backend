package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
	"github.com/spf13/viper"
	"github.com/zi-yang-zhang/audible-silence-backend/api"
)

func main() {
	app := iris.Default()
	config := loadConfiguration()
	dbConfig := config.GetStringMapString("db")
	serverConfig := config.GetStringMapString("server")
	db, err := gorm.Open("mysql", dbConfig["connection"])
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Error connecting to DB: %s", err))
	}
	defer db.Close()
	api.InitAPI(app, db, config)
	app.Run(iris.Addr(":" + serverConfig["port"]))
}

func loadConfiguration() *viper.Viper {
	config := viper.New()

	config.SetConfigName("conf")
	config.AddConfigPath(".")
	err := config.ReadInConfig() // Find and read the config file
	if err != nil {
		// Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s", err))

	}
	return config

}
