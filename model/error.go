package model

import (
	"encoding/json"
	"reflect"
)

//ValidationError struct for error response
type ValidationError struct {
	Header []string `json:"header,omitempty"`
	Body   []string `json:"body,omitempty"`
}

func (e ValidationError) Error() string {
	json, _ := json.Marshal(e)
	return string(json)
}

func (e *ValidationError) Is(target error) bool {
	return reflect.TypeOf(target) == reflect.TypeOf(e)
}

//UnauthorizedError represent auth erro
type UnauthorizedError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func (e *UnauthorizedError) Is(target error) bool {
	return reflect.TypeOf(target) == reflect.TypeOf(e)
}

//ErrorMsgs model with all error text messages
var ErrorMsgs  = map[string]string{
	"userNotExists": "User not exists or login with password not correct",
	"emptyEmail": "User.Email: zero value",
	"emptyPassword": "User.Password: zero value",
	"emptyUsername":  "User.Username: zero value",
	"bodyNotEmpty": "can't be empty",
}