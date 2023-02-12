package service

import (
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/rest_err"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"github.com/andrade-felipe-dev/jwt-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(domain.UserDomainInterface) (
		domain.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, domain.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*domain.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
