package service

import (
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/logger"
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/rest_err"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain domain.UserDomainInterface) (domain.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	return userDomainRepository, nil
}
