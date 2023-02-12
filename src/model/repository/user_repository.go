package repository

import (
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/rest_err"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type UserRepository interface {
	CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr)
}

type userRepository struct {
	databaseConnection *mongo.Database
}
