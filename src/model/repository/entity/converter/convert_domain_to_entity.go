package converter

import (
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"github.com/andrade-felipe-dev/jwt-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain domain.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email: domain.GetEmail(),
		Senha: domain.GetSenha(),
		Nome:  domain.GetNome(),
		Tipo:  domain.GetTipo(),
	}
}
