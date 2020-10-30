package main

import (
	srv "github.com/dmitry-dms/rest-gin"
	hs "github.com/dmitry-dms/rest-gin/pkg/handler"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
	"github.com/dmitry-dms/rest-gin/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialising config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := hs.NewHandler(services)

	server := new(srv.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
