package model

import (
	"fmt"
	"ql/crud/src/configuration/opps"
)

func (ud *UserDomain) CreateUser() *opps.Opps {
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}