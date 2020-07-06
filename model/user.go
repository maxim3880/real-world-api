package model

//User represent table model of users
type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//RequestUserModel request model to authentication or sign up user
type RequestUserModel struct {
	User User `json:"user"`
}

//ResponseUserModel request model to sign up user
type ResponseUserModel struct {
	User User `json:"user"`
}
