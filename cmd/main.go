package main

import (
	todo "github.com/Goganad/TodoList-REST-API"
	"github.com/Goganad/TodoList-REST-API/handlers"
	"github.com/Goganad/TodoList-REST-API/repository"
	"github.com/Goganad/TodoList-REST-API/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main(){
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Error initializing DB: %s", err.Error())
	}

	repos:=repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server: %s", err.Error())
	}
}


//igornadenenko$ docker run --name=todo-db -e POSTGRES_PASSWORD='butterfly3000' -d --rm postgres
//migrate -path ./schema -database 'postgres://postgres:butterfly3000@localhost5432/postgres?sslmode-disable' up
