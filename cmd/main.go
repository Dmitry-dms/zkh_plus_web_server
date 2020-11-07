package main

import (
	"context"
	srv "github.com/dmitry-dms/rest-gin"
	"github.com/dmitry-dms/rest-gin/pkg/handler"
	"github.com/dmitry-dms/rest-gin/pkg/repository"
	"github.com/dmitry-dms/rest-gin/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
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
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		DBName:   viper.GetString("db.dbname"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(srv.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error while running: %s", err.Error())
		}
	}()
	logrus.Print("ZKH plus server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) // сигналы в unix системах
	<-quit

	logrus.Print("ZKH plus server are shutting down")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred while shutting down: %s", err.Error())
	}
	if //goland:noinspection ALL
	err := db.Close(); err != nil {
		logrus.Errorf("error occurred while closing db: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
