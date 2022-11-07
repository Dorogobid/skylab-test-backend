package main

import (
	"fmt"
	"os"

	"github.com/kardianos/osext"
	"github.com/spf13/viper"
)

func initConfig() error {
	path, err := osext.ExecutableFolder()
	if err != nil {
		panic("Error initializing configs")
	}
	fmt.Println(path)
	viper.AddConfigPath(path + "/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getConfig() *DBConfig {
	host, isHostPresent := os.LookupEnv("DB_HOST")
	if !isHostPresent {
		host = viper.GetString("db.host")
	}

	return &DBConfig{
		Host:     host,
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		Port:     viper.GetString("db.port"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}