package model

import (
	"database/sql"
	"encoding/json"
)

//User represent table model of users
type User struct {
	ID       int          `json:"-" db:"id"`
	Email    string       `json:"email" db:"email"`
	Username string       `json:"username" db:"name"`
	Password string       `json:"-" db:"password"`
	Token    string       `json:"token"`
	Bio      dbNullString `json:"bio,string" db:"bio"`
	Image    dbNullString `json:"image,string" db:"image"`
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

//UpdateUserRequestModel request model to update user
type UpdateUserRequestModel struct {
	User updateFields `json:"user"`
}

type updateFields struct {
	Email    string `json:"email" structs:"email,omitempty"`
	Password string `json:"password" structs:"password,omitempty"`
	Username string `json:"username" structs:"name,omitempty"`
	Bio      string `json:"bio" structs:"bio,omitempty"`
	Image    string `json:"image" structs:"image,omitempty"`
}

type dbNullString struct {
	sql.NullString
}

func (s dbNullString) MarshalJSON() ([]byte, error) {
	if s.Valid {
		return json.Marshal(s.String)
	}
	return []byte(`null`), nil
}
