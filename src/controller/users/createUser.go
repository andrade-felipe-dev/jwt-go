package users

import (
	"net/http"

	"github.com/andrade-felipe-dev/jwt-go/src/configuration/logger"
	"github.com/andrade-felipe-dev/jwt-go/src/configuration/validation"
	"github.com/andrade-felipe-dev/jwt-go/src/controller/model/request"
	"github.com/andrade-felipe-dev/jwt-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar o usuário", err, zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)
		c.JSON(restErr.Code, restErr)
		return
	}

	logger.Info("User created successfully", zap.String("journey", "createUser"))

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Nome:  userRequest.Nome,
		Tipo:  userRequest.Tipo,
	}

	c.JSON(http.StatusOK, response)
}
