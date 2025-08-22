package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port          int
	MysqlAccount  string
	MysqlPassword string
	MysqlAddress  string
	MysqlDBName   string
}

func GetConfig() Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Printf("Warning: config PORT=\"%s\", which is not a integer", os.Getenv("PORT"))
		port = 80
	}
	result := Config{
		Port:          port,
		MysqlAccount:  os.Getenv("MYSQL_ACCOUNT"),
		MysqlPassword: os.Getenv("MYSQL_PWD"),
		MysqlAddress:  os.Getenv("MYSQL_ADDR"),
		MysqlDBName:   os.Getenv("MYSQL_DB_NAME")}
	return result
}
