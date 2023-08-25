package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ConfigDatabase struct {
	Driver     string
	dbUsername string
	dbPassword string
	dbHost     string
	dbPort     string
	dbName     string
}

func (conf *ConfigDatabase) NewConfigDatabase() {
	err := godotenv.Load("conf/config.env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	conf.Driver = os.Getenv("dbDriver")
	conf.dbUsername = os.Getenv("dbUsername")
	conf.dbPassword = os.Getenv("dbPassword")
	conf.dbHost = os.Getenv("dbHost")
	conf.dbPort = os.Getenv("dbPort")
	conf.dbName = os.Getenv("dbName")
	fmt.Println(conf)
}

func (conf ConfigDatabase) GetURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.dbUsername, conf.dbPassword, conf.dbHost, conf.dbPort, conf.dbName)
}
