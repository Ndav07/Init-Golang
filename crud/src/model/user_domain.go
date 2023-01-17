package model

import (
	"crypto/md5"
	"encoding/hex"
	"ql/crud/src/configuration/opps"
)

type UserDomain struct {
	Email string
	Password string 
	Name string 
	Age int8 
}

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &UserDomain{
		Email: email,
		Password: password,
		Name: name,
		Age: age,
	}
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil)) 
}

type UserDomainInterface interface {
	CreateUser() *opps.Opps
	UpdateUser(string) *opps.Opps
	FindUser(string) (*UserDomain, *opps.Opps)
	DeleteUser(string) *opps.Opps
}

