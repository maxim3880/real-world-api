package service

import "gopkg.in/validator.v2"

var (
	//LoginValidator use for validate model in login handle
	LoginValidator = validator.NewValidator()
	//RegisterValidator use for validate model in registration handle
	RegisterValidator = validator.NewValidator()
)

//SetValidatorTags init all validation tags
func SetValidatorTags() {
	LoginValidator.SetTag("login")
	RegisterValidator.SetTag("signup")
}
