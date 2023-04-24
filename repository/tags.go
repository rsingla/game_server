package repository

import "github.com/rsingla/game_server/model"

type TagsRepository interface {
	FindAll() ([]*model.Tags, error)
	FindById(id int) (*model.Tags, error)
	FindBySlug(slug string) (*model.Tags, error)
	Save(tags *model.Tags) (*model.Tags, error)
	Update(tags *model.Tags) (*model.Tags, error)
	DeleteTags(id int) error
}
