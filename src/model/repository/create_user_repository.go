package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/andrade-felipe-dev/jwt-go/src/configuration/logger"
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/rest_err"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"github.com/andrade-felipe-dev/jwt-go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGO_USER_DB = "MONGO_USER_DB"
)

func (ur *userRepository) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init create user repository")
	collection_name := os.Getenv(MONGO_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	entityUser := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), entityUser)
	fmt.Println("pós insert")

	if err == nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	entityUser.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*entityUser), nil
}
