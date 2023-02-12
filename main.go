package main

import (
	"context"
	"log"

	"github.com/andrade-felipe-dev/jwt-go/src/configuration/database/mongodb"
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/logger"
	"github.com/andrade-felipe-dev/jwt-go/src/controller/routes"
	"github.com/andrade-felipe-dev/jwt-go/src/controller/users"
	"github.com/andrade-felipe-dev/jwt-go/src/model/repository"
	"github.com/andrade-felipe-dev/jwt-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Começando o projeto...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := users.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
