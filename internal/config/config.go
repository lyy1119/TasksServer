package config

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	Port          int
	MysqlAccount  string
	MysqlPassword string
	MysqlAddress  string
	MysqlDBName   string
}

func GetConfig() (Config, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Warn(fmt.Sprintf("config PORT=\"%s\", which is not a integer. Use 80 as default", os.Getenv("PORT")))
		port = 80
	}
	result := Config{
		Port:          port,
		MysqlAccount:  os.Getenv("MYSQL_ACCOUNT"),
		MysqlPassword: os.Getenv("MYSQL_PWD"),
		MysqlAddress:  os.Getenv("MYSQL_ADDR"),
		MysqlDBName:   os.Getenv("MYSQL_DATABASE")}
	return result, err
}
