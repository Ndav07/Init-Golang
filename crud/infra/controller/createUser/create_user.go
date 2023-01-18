package createUser

import (
	"fmt"
	"net/http"
	"ql/crud/application/usecase/createUser"
	"ql/crud/config/logger"
	"ql/crud/config/validation"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser Controller",
		zap.String("jorney", "createUser"),
	)
	var userInput createUser.CreateUserInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		logger.Error("Error trying to validate user info", err)
		opps := validation.ValidateUserError(err)
		c.JSON(opps.Code, opps)
		return
	}

	
	user, err := createUser.CreateUser(userInput)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	fmt.Println(user)

	c.String(http.StatusOK, "")
}