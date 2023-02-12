package converter

import (
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"github.com/andrade-felipe-dev/jwt-go/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) domain.UserDomainInterface {
	domain := domain.NewUserDomain(
		entity.Email,
		entity.Senha,
		entity.Nome,
		entity.Tipo,
	)

	domain.SetID(entity.ID.Hex())
	return domain
}
