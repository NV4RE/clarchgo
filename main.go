package main

import (
	fiber "clarchgo/controller/fiber"
	"clarchgo/repository/auth"
	authUseCase "clarchgo/use-case/auth"
	"fmt"
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	App struct {
		Name        string `env:"APP_NAME" envDefault:"clarchgo"`
		Version     string `env:"APP_VERSION" envDefault:"0.0.0"`
		Environment string `env:"APP_ENVIRONMENT" envDefault:"local"` // e.g. local, development, staging, uat, production
	}
	HttpServer struct {
		Host string `env:"HTTP_SERVER_HOST" envDefault:"0.0.0.0"`
		Port int    `env:"HTTP_SERVER_PORT" envDefault:"8800"`
	}
	MongoURI    string `env:"MONGODB_URI" envDefault:"mongodb://localhost:27017/demo"`
	RedisServer string `env:"REDIS_URI" envDefault:"redis://localhost:16379/0"`
}

func main() {
	// Load config
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln(err)
	}

	// Set-up repository e.g. database, service connection
	authRepo, err := auth.NewMongo(cfg.MongoURI, fmt.Sprintf("%s-%s", cfg.App.Environment, cfg.App.Name))
	if err != nil {
		log.Fatalln(err)
	}

	// Set-up use-case e.g. business
	authUC := authUseCase.NewUseCase(authRepo)

	// Set-up controller/interface
	fs := fiber.NewFiberServer(authUC)
	fs.Serve(fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port))
}
