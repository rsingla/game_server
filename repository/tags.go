package repository

type TagsRepository interface {
	FindAll() ([]*Tags, error)
	FindById(id int) (*Tags, error)
	FindBySlug(slug string) (*Tags, error)
	Store(tags *Tags) (*Tags, error)
	Update(tags *Tags) (*Tags, error)
	Delete(tags *Tags) error
}
