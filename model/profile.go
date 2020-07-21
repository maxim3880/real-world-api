package model

//Profile represent model of user
type Profile struct {
	ID        int          `json:"-" db:"id"`
	Username  string       `json:"username" db:"name"`
	Bio       dbNullString `json:"bio,string" db:"bio"`
	Image     dbNullString `json:"image,string" db:"image"`
	Following bool         `json:"following"`
}

type ProfileResponse struct {
	Profile Profile `json:"profile"`
}
