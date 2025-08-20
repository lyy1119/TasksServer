package config

import (
	"os"
)

type Config struct {
	mysqlAccount  string
	mysqlPassword string
}

func GetConfig() Config {
	result := Config{
		mysqlAccount:  os.Getenv("MYSQL_ACCOUNT"),
		mysqlPassword: os.Getenv("MYSQL_PWD")}

	return result
}
