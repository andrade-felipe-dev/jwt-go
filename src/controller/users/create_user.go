package users

import (
	"net/http"

	"github.com/andrade-felipe-dev/jwt-go/src/configuration/logger"
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/validation"
	"github.com/andrade-felipe-dev/jwt-go/src/controller/model/request"
	"github.com/andrade-felipe-dev/jwt-go/src/model/domain"
	"github.com/andrade-felipe-dev/jwt-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	userModel domain.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init a CreateUser controller", zap.String("journey", "createUser"))

	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar o usuário", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	domainModel := domain.NewUserDomain(
		userRequest.Email,
		userRequest.Senha,
		userRequest.Nome,
		userRequest.Tipo,
	)
	domainResult, err := uc.service.CreateUser(domainModel)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
