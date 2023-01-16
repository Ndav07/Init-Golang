package controller

import (
	"fmt"
	"ql/crud/src/configuration/validation"
	"ql/crud/src/model/request"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		opps := validation.ValidateUserError(err)
		c.JSON(opps.Code, opps)
		return
	}

	fmt.Println(userRequest)
}