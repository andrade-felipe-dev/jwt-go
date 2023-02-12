package view

import (
	"github.com/andrade-felipe-dev/jwt-go/src/controller/model/response"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
)

func ConvertDomainToResponse(
	userDomain domain.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Nome:  userDomain.GetNome(),
		Tipo:  userDomain.GetTipo(),
	}
}
