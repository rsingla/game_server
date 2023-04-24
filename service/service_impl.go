package service

import (
	"github.com/go-playground/validator"
	"github.com/rsingla/game_server/helper"
	"github.com/rsingla/game_server/model"
	"github.com/rsingla/game_server/repository"
	"github.com/rsingla/game_server/request"
)

type TagsService struct {
	TagsRepository repository.TagsRepository
	validate       *validator.Validate
}

func (t *TagsService) Save(tagsReq *request.TagsReq) (*model.Tags, error) {

	//err := t.validate.Struct(&tagsReq)
	//if err != nil {
	//	log.Println(err)
	//	log.Println(err.Error())
	//	return nil, err
	//}

	tagsImpl := repository.TagsImpl{Db: helper.ConnectDB()}

	//Convert tags to model.Tags
	tagsModel := model.Tags{Name: tagsReq.Name, Slug: tagsReq.Slug, Description: tagsReq.Description, Status: tagsReq.Status, Type: tagsReq.Type, Category: tagsReq.Category}

	modelTags, err := tagsImpl.Save(&tagsModel)

	if err != nil {
		return nil, err
	}

	return modelTags, nil

}

func (t *TagsService) FindAll() ([]*model.Tags, error) {
	err := t.validate.Struct(t)
	if err != nil {
		return nil, err
	}

	return t.TagsRepository.FindAll()

}

func (t *TagsService) FindById(id int) (*model.Tags, error) {
	err := t.validate.Struct(t)
	if err != nil {
		return nil, err
	}
	return t.TagsRepository.FindById(id)
}

func (t *TagsService) FindBySlug(slug string) (*model.Tags, error) {
	err := t.validate.Struct(t)
	if err != nil {
		return nil, err
	}
	return t.TagsRepository.FindBySlug(slug)
}

func (t *TagsService) Update(tagsReq *request.TagsReq) (*model.Tags, error) {
	err := t.validate.Struct(t)
	if err != nil {
		return nil, err
	}
	//Convert tags to model.Tags
	tagsModel := model.Tags{Name: tagsReq.Name, Slug: tagsReq.Slug, Description: tagsReq.Description, Status: tagsReq.Status, Type: tagsReq.Type, Category: tagsReq.Category}

	return t.TagsRepository.Update(&tagsModel)
}

func (t *TagsService) Delete(id int) error {

	tagsImpl := repository.TagsImpl{Db: helper.ConnectDB()}

	err := tagsImpl.DeleteTags(id)

	if err != nil {
		return err
	}

	return nil
}
