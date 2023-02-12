package routes

import (
	"github.com/andrade-felipe-dev/jwt-go/src/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController users.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
