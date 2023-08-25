package conf

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config_MailService struct {
	From     string
	password string
	Host     string
	Port     uint
}

func (conf *Config_MailService) NewConfig_MailService() {
	err := godotenv.Load("conf/config.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	conf.From = "kiettranuit@gmail.com"
	conf.password = os.Getenv("passEmail")
	conf.Host = "smtp.gmail.com"
	conf.Port = 587
}

func (conf *Config_MailService) GetPassword() string {
	return conf.password
}

func (conf Config_MailService) URL() string {
	return fmt.Sprintf("%s:%d", conf.Host, conf.Port)
}
