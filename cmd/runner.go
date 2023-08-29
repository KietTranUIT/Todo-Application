package main

import (
	"log"
	"user-service/conf"
	"user-service/internal/controller"
	"user-service/internal/core/infra/repository"
	"user-service/internal/core/server"
	"user-service/internal/core/service"
)

func main() {
	log.Println("Server is running ...")
	var configDatabase conf.ConfigDatabase
	configDatabase.NewConfigDatabase()
	db, err := repository.NewDB(configDatabase)

	if err != nil {
		log.Println(err.Error())
	}

	userRepo := repository.NewUserRepo(db)

	var configMailService conf.Config_MailService
	configMailService.NewConfig_MailService()
	mail := service.NewMailService(configMailService)
	userService := service.NewUserService(userRepo, mail)

	userController := controller.NewUserController(userService)

	userController.Router()

	httpServer := server.NewHTTPServer(conf.ConfigServer{
		Host: "127.0.0.1",
		Port: 9000,
	}, userController.Mux)

	httpServer.Start()
}
