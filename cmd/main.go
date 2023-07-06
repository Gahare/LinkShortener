package main

import (
	"LinkShortener"
	"LinkShortener/pkg/handler"
	"LinkShortener/pkg/repository"
	"LinkShortener/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	db, errdb := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		Password: viper.GetString("db.Password"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})
	if errdb != nil {
		logrus.Fatalf("failed to initialize db: %s", errdb.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(LinkShortener.Server)
	go func() {
		errhttp := srv.Run(viper.GetString("serverport"), handlers.InitRoutes())
		if errhttp != nil {
			logrus.Fatalf("error occured while running http server: %s", errhttp.Error())
		}
	}()
	logrus.Println("LinkShortener started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
