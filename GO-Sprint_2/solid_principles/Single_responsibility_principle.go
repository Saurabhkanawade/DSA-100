package main

import (
	"errors"
	"fmt"
)

type UserNOTSRP struct {
	FirstName string
	LastName  string
}

func (u *UserNOTSRP) GetFullName() string {
	return u.FirstName + " " + u.LastName
}

func (u *UserNOTSRP) Save() error {
	// Save user to the database
	// ...
	return errors.New("error")
}

//adhering to srp

type User struct {
	name  string
	email string
}

func (u *User) GetUser() {
	fmt.Println("The user is :", u.email)
}

type UserRepository struct {
	// Database connection or other storage-related fields
}

func (u *UserRepository) SaveUser() {
	fmt.Println("saving the user", u)
}

func main() {
	user := User{
		name:  "Saurabh Kanawade",
		email: "saurabhkanawade30@gmail.com",
	}
	user.GetUser()
	fmt.Println(user)
	ur := UserRepository{}
	ur.SaveUser()

	u := UserNOTSRP{
		FirstName: "Not srp",
		LastName:  "Not srp",
	}
	err := u.Save()
	if err != nil {
		return
	}
	u.GetFullName()
}
