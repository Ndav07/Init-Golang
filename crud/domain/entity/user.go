package entity

type User struct {
	email string
	password string 
	name string 
	age int8 
}

func NewUser(email, password, name string, age int8) *User {
	return &User{
		email: email,
		password: password,
		name: name,
		age: age,
	}
}

func (ud *User) GetEmail() string {
	return ud.email
}

func (ud *User) GetPassword() string {
	return ud.password
}

func (ud *User) GetName() string {
	return ud.name
}

func (ud *User) GetAge() int8 {
	return ud.age
}