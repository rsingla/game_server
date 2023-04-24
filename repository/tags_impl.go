package repository

import (
	"github.com/rsingla/game_server/model"
	"gorm.io/gorm"
)

type TagsImpl struct {
	Db *gorm.DB
}

func (t *TagsImpl) FindAll() ([]*model.Tags, error) {
	var tags []*model.Tags
	err := t.Db.Find(&tags).Error
	return tags, err
}

func (t *TagsImpl) FindById(id int) (*model.Tags, error) {
	var tags model.Tags
	err := t.Db.First(&tags, id).Error
	return &tags, err
}

func (t *TagsImpl) FindBySlug(slug string) (*model.Tags, error) {
	var tags model.Tags
	err := t.Db.Where("slug = ?", slug).First(&tags).Error
	return &tags, err
}

func (t *TagsImpl) Save(tags *model.Tags) (*model.Tags, error) {
	err := t.Db.Create(tags).Error
	return tags, err
}

func (t *TagsImpl) Update(tags *model.Tags) (*model.Tags, error) {
	err := t.Db.Save(tags).Error
	return tags, err
}

func (t *TagsImpl) Delete(tags *model.Tags) error {
	err := t.Db.Delete(tags).Error
	return err
}
