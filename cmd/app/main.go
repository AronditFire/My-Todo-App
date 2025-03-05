package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AronditFire/TODO-APP/config"
	"github.com/AronditFire/TODO-APP/internal/database"
	"github.com/AronditFire/TODO-APP/internal/entities"
	"github.com/AronditFire/TODO-APP/internal/handlers"
	"github.com/AronditFire/TODO-APP/internal/repository"
	"github.com/AronditFire/TODO-APP/internal/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	//Load Env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error in loading env values: %s", err.Error())
		return
	}
	//Load cfg
	if err := config.ReadConfigFile(); err != nil {
		log.Fatalf("error in reading config file : %s", err.Error())
		return
	}

	db, err := database.ConnectToDB(database.Config{
		Host:     viper.GetString("db.host"),
		DBname:   viper.GetString("db.dbname"),
		Username: viper.GetString("db.username"),
		Port:     viper.GetString("db.port"),
		Password: os.Getenv("DB_PASSWORD"),
		Sslmode:  viper.GetString("db.sslmode"),
		TimeZone: viper.GetString("db.timezone"),
	})
	if err != nil {
		log.Fatalf("connection to db error: %s", err.Error())
		return
	}

	repos := repository.CreateRepository(db)
	service := service.NewService(repos)
	handler := handlers.NewHander(service)

	sv := new(entities.Service)

	go func() {
		if err := sv.RunServer(viper.GetString("port"), handler.InitRouters()); err != nil {
			log.Fatalf("error of server running: %s", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := sv.Shutdown(); err != nil {
		log.Fatalf("shutdown server error: %s", err.Error())
	}

	database.CloseDB(db)
}
