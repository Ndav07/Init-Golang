package createUser

import (
	"crypto/md5"
	"encoding/hex"
	"ql/crud/config/opps"
	"ql/crud/domain/entity"
)

func CreateUser(userInput CreateUserInput) (entity.User ,*opps.Opps) {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userInput.Password))
	newPassword := hex.EncodeToString(hash.Sum(nil)) 

	user := entity.NewUser(userInput.Email, newPassword, userInput.Name, userInput.Age)

	return *user, nil
}