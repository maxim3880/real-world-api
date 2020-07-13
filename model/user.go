package model

//User represent table model of users
type User struct {
	ID       int    `json:"-" db:"id"`
	Email    string `json:"email" db:"email"`
	Username string `json:"username" db:"name"`
	Password string `json:"-" db:"password"`
	Token    string `json:"token"`
	Bio      string `json:"bio" db:"bio"`
	Image    string `json:"image" db:"image"`
}

//AuthRequestModel request model to login or sign up user
type AuthRequestModel struct {
	User authUserField `json:"user"`
}

type authUserField struct {
	Email    string `json:"email" login:"nonzero" signup:"nonzero"`
	Password string `json:"password" login:"nonzero" signup:"nonzero"`
	Username string `json:"username" signup:"nonzero"`
}

//ResponseUserModel request model to sign up user
type ResponseUserModel struct {
	User *User `json:"user"`
}
