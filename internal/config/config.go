package config

import (
	"os"
)

type Config struct {
	MysqlAccount  string
	MysqlPassword string
	MysqlAddress  string
	MysqlDBName   string
}

func GetConfig() Config {
	result := Config{
		MysqlAccount:  os.Getenv("MYSQL_ACCOUNT"),
		MysqlPassword: os.Getenv("MYSQL_PWD"),
		MysqlAddress:  os.Getenv("MYSQL_ADDR"),
		MysqlDBName:   os.Getenv("MYSQL_DB_NAME")}
	return result
}
