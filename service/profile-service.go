package service

import (
	"github.com/maxim3880/real-world-api/data"
	"github.com/maxim3880/real-world-api/model"
)

//ProfileService represent repository for profiles
type ProfileService struct {
	data.Store
}

//GetProfileByUsername return profile model by username
func (s *ProfileService) GetProfileByUsername(username string) (res model.Profile, err error) {
	err = s.Store.Get(&res, "SELECT id, name, bio, image FROM users WHERE name = $1", username)
	return res, err
}

//CreateProfileFollow new follow on profile
func (s *ProfileService) CreateProfileFollow(username string, userID int) (res model.Profile, err error) {
	res, err = s.GetProfileByUsername(username)
	if err != nil {
		return
	}
	s.Store.Insert("INSERT INTO user_follows (follow_user_id, user_id) VALUES($1, $2)", res.ID, userID)
	res.Following = true
	return res, err
}

//DeleteProfileFollow return profile model by username
func (s *ProfileService) DeleteProfileFollow(username string, userID int) (res model.Profile, err error) {
	res, err = s.GetProfileByUsername(username)
	if err != nil {
		return
	}
	_, err = s.Store.Delete("DELETE FROM user_follows WHERE follow_user_id = $1 AND user_id = $2", res.ID, userID)
	return res, err
}
