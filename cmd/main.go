package main

import (
	handlerPKG "file_work/internal/handler"
	"file_work/internal/repository"
	"file_work/internal/server"
	servicePKG "file_work/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"

	_ "github.com/lib/pq"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	srv := new(server.Server)

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DbName:   viper.GetString("db.dbname"),
		SSlmode:  viper.GetString("db.sslmod"),
	})
	if err != nil {
		log.Fatalf("failed to initializing database: %s ", err.Error())
		return
	}

	defer db.Close()

	repos := repository.NewRepository(db)
	service := servicePKG.NewService(repos)
	handler := handlerPKG.NewHandler(service)

	go func() {
		if err := srv.Run("9009", handler.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
			return
		}
	}()

	logrus.Println("server starting...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("configs")

	return viper.ReadInConfig()
}
