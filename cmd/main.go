package main

import (
	"context"
	todo "github.com/Goganad/TodoList-REST-API"
	"github.com/Goganad/TodoList-REST-API/handlers"
	"github.com/Goganad/TodoList-REST-API/repository"
	"github.com/Goganad/TodoList-REST-API/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error initializing DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error occured while running server: %s", err.Error())
		}
	}()
	log.Printf("Server has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server is shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error on shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Error on closing database connection: %s", err.Error())
	}
}

//igornadenenko$ docker run --name=todo-db -e POSTGRES_PASSWORD='butterfly3000' -d --rm postgres
//migrate -path ./schema -database 'postgres://postgres:butterfly3000@localhost:5436/postgres?sslmode=disable' up
