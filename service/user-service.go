package service

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/structs"
	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/model"
)

//UserService represent user logic
type UserService struct {
	data.Store
}

//LoginUser method for login user by input model
func (us *UserService) LoginUser(data model.AuthRequestModel) (user model.User, err error) {
	user = model.User{}
	user.Email = data.User.Email
	user.Password = data.User.Password
	err = us.Store.Get(&user, "SELECT * FROM users WHERE email=$1 AND password=$2", user.Email, user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New(model.ErrorMsgs["userNotExists"])
		}
		return user, err
	}
	user.Token = GenerateJwtTocken(user.Email, user.ID)
	return user, err
}

//RegisterUser method for login user by input model
func (us *UserService) RegisterUser(data model.AuthRequestModel) (user model.User, err error) {
	user = model.User{}
	err = us.Store.Get(&user, "SELECT * FROM users WHERE email=$1 OR name=$2", data.User.Email, data.User.Username)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println(err)
		return user, err
	} else if err == sql.ErrNoRows {
		user, err = us.AddUser(data)
		user.Token = GenerateJwtTocken(user.Email, user.ID)
	} else {
		err = errors.New("User already registered with same email or username")
	}
	return user, err
}

//AddUser represent adding one user to db
func (us *UserService) AddUser(data model.AuthRequestModel) (user model.User, err error) {
	id := us.Store.Insert("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", data.User.Username, data.User.Email, data.User.Password)
	user, err = us.GetUserByID(id)
	return
}

//GetUserByID get user model by user ID from db
func (us *UserService) GetUserByID(id int) (user model.User, err error) {
	err = us.Store.Get(&user, "SELECT * FROM users where id=$1", id)
	return user, err
}

//UpdateUser update user data table in db
func (us *UserService) UpdateUser(data model.UpdateUserRequestModel, id int) (user model.User, err error) {
	val := structs.Map(data.User)
	columns := []string{}
	for col := range val {
		columns = append(columns, col+"=:"+col)
	}
	conCol := strings.Join(columns, ", ")
	val["id"] = id
	_, err = us.Store.Update("UPDATE users SET "+conCol+" WHERE id=:id", val)
	if err != nil {
		return user, err
	}
	user, err = us.GetUserByID(id)
	if err != nil {
		return
	}
	return
}
