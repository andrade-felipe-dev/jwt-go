package service

import (
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/rest_err"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
)

func (*userDomainService) UpdateUser(userId string, userDomain domain.UserDomainInterface) *rest_err.RestErr {
	return nil
}
