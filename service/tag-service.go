package service

import (
	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/model"
)

//TagService represent all methods of tags
type TagService struct {
	data.Store
}

//GetAll return all tags from db
func (t *TagService) GetAll() []model.Tag {
	tags := make([]model.Tag, 0)
	t.Store.Select(&tags, "Select * from tags")
	return tags
}
