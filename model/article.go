package model

import "time"

//Article represent table articles in db
type Article struct {
	ID          int    `json:"-" db:"id"`
	Slug        string `json:"slug" db:"slug"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Body        string `json:"body" db:"body"`
	CreatedAt   string `json:"createdAt" db:"createdAt"`
	UpdatedAt   string `json:"updatedAt" db:"updatedAt"`
	AuthorID    int    `json:"-" db:"author_id"`
}

//ResponseArticle represent 1 article model in response
type ResponseArticle struct {
	Slug           string    `json:"slug" db:"slug"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	Body           string    `json:"body" db:"body"`
	CreatedAt      time.Time `json:"createdAt" db:"createdat"`
	UpdatedAt      time.Time `json:"updatedAt" db:"updatedat"`
	Author         Profile   `json:"author"`
	Favorited      bool      `json:"favorited" db:"favorited"`
	FavoritesCount int       `json:"favoritesCount" db:"favoritesCount"`
	TagList        []string  `json:"tagList"`
}

//MultiArticleResponse represent model for more than one arcticle response
type MultiArticleResponse struct {
	Articles      []ResponseArticle `json:"articles"`
	ArticlesCount int               `json:"articlesCount"`
}

//SingleArticleResponse represent model for one arcticle response
type SingleArticleResponse struct {
	Articles ResponseArticle `json:"articles"`
}
