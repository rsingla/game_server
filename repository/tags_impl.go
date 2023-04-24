package repository

import (
	"errors"
	"log"

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
	tags := model.Tags{}

	err := t.Db.First(&tags, id).Error

	if err != nil {
		return &tags, errors.New("tag not found")
	}

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

func (t *TagsImpl) DeleteTags(id int) error {
	log.Println("Delete : ", id)
	var tags model.Tags

	err := t.Db.First(&tags, id).Error

	if err != nil {
		return errors.New("tag not found")
	}

	log.Println(tags)

	err = t.Db.Delete(tags).Error

	return err
}
