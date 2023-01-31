package routes

import (
	"github.com/andrade-felipe-dev/jwt-go/src/controller/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", users.FindUserById)
	r.GET("/getUserByEmail", users.FindUserByEmail)
	r.POST("/createUser", users.CreateUser)
	r.PUT("/updateUser/:userId", users.UpdateUser)
	r.DELETE("/deleteUser/:userId", users.DeleteUser)
}
