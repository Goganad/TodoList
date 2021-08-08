package main

import (
	"context"
	"github.com/Goganad/TodoList-REST-API"
	"github.com/Goganad/TodoList-REST-API/pkg/handlers"
	"github.com/Goganad/TodoList-REST-API/pkg/repository"
	"github.com/Goganad/TodoList-REST-API/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const configPath = "config"

func initConfig() error {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configPath)
	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("Initializing DB: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("While running server: %s", err.Error())
		}
	}()
	log.Printf("Server has started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server is shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("On shutting down server: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("On closing database connection: %s", err.Error())
	}
}
