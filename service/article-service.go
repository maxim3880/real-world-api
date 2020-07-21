package service

import (
	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/model"
)

//ArticleService represent all methods of tags
type ArticleService struct {
	data.Store
}

//GetAll return all tags from db
func (a *ArticleService) GetAll() []model.ResponseArticle {
	articles := make([]model.ResponseArticle, 0)
	sql := `SELECT
				articles.slug, 
				articles.title, 
				articles.description, 
				articles.body, 
				articles.createdAt, 
				articles.updatedAt,
				users.name  "author.name",
				users.bio  "author.bio",
				users.image  "author.image",
				EXISTS (SELECT NULL FROM user_favorite_articles WHERE articles.id = user_favorite_articles.article_id and user_id = 1) "favorited",
				(SELECT count(1) FROM user_favorite_articles WHERE articles.id = user_favorite_articles.article_id) "favoritesCount"
			FROM articles 
			INNER JOIN users  ON articles.author_id = users.id`
	err := a.Store.Select(&articles, sql)
	if err != nil {
		return nil
	}
	// res := a.gerateResponse(articles)
	return articles
}

func (a *ArticleService) gerateResponse(data []model.Article) []model.ResponseArticle {
	res := []model.ResponseArticle{}
	return res
}
