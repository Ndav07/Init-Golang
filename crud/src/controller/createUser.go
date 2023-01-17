package controller

import (
	"net/http"
	"ql/crud/src/configuration/logger"
	"ql/crud/src/configuration/validation"
	"ql/crud/src/model"
	"ql/crud/src/model/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("jorney", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err)
		opps := validation.ValidateUserError(err)
		c.JSON(opps.Code, opps)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.String(http.StatusOK, "")
}