package routes

import (
	"ql/crud/infra/controller/createUser"
	"ql/crud/infra/controller/deleteUser"
	"ql/crud/infra/controller/findUserByEmail"
	"ql/crud/infra/controller/findUserById"
	"ql/crud/infra/controller/updateUser"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", findUserById.FindUserById)
	r.GET("/getUserByEmail/:userEmail", findUserByEmail.FindUserByEmail)
	r.POST("/createUser", createUser.CreateUser)
	r.PUT("/updateUser/:userId", updateUser.UpdateUser)
	r.DELETE("/deleteUser/:userId", deleteUser.DeleteUser)
}