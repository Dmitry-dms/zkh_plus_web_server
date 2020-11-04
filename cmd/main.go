package main

import (
	//gorm "github.com/jinzhu/gorm"
	//"database/sql"
	srv "github.com/dmitry-dms/rest-gin"
	"github.com/dmitry-dms/rest-gin/pkg/handler"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
	"github.com/dmitry-dms/rest-gin/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	//"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialising config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "fullstack-postgres", //viper.GetString("db.host"),
		Port:     "5436",               //viper.GetString("db.port"),
		DBName:   "postgres",           //viper.GetString("db.dbname"),
		Username: "postgres",           //viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  "disable", //viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(srv.Server)

	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error while running: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
