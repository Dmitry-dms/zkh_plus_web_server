package main

import (
	"context"
	srv "github.com/Dmitry-dms/zkh-plus"
	"github.com/Dmitry-dms/zkh-plus/pkg/handler"
	"github.com/Dmitry-dms/zkh-plus/pkg/repository"
	"github.com/Dmitry-dms/zkh-plus/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env: %s", err.Error())
	}
	// db, err := repository.NewPostgresDB(repository.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	DBName:   os.Getenv("DB_NAME"),
	// 	Username: os.Getenv("DB_USERNAME"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// })
	// if err != nil {
	// 	logrus.Fatalf("failed to initialize db: %s", err.Error())
	// }

	//repos := repository.NewRepository(db)
	mongoDatabase, err := repository.NewMongoDB(context.TODO(), "mongodb://localhost:7000/?readPreference=primary&appname=MongoDB%20Compass&ssl=false", "zkh")
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	mongo := repository.NewMongoRepository(mongoDatabase)
	services := service.NewService(mongo)
	handlers := handler.NewHandler(services)

	server := new(srv.Server)
	go func() {
		if err := server.Run(os.Getenv("API_PORT"), handlers.InitRoutes()); err != nil {
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
	// if //goland:noinspection ALL
	// err := db.Close(); err != nil {
	// 	logrus.Errorf("error occurred while closing db: %s", err.Error())
	// }
}
