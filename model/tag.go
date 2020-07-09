package model

//Tag represent tag db model
type Tag struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

//ResponseTagModel represent response model of tag controller
type ResponseTagModel struct {
	Tags []string `json:"tags"`
}
