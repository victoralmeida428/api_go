package repository

import "api/src/model"

type IRepository[T any] interface {
	FindAll() ([]T, error)
	FindById(id int) (T, error)
	FindByField(field string, value interface{}) ([]T, error)
	Update(T) (T, error)
	Insert(T) (T, error)
	Delete(T) (T, error)
}

type Repository struct {
	model *model.Model
	Grupo IRepository[model.Grupo]
}

func New(md *model.Model) *Repository {
	return &Repository{
		model: md,
		Grupo: &Grupo{db: md.Grupo.DB},
	}
}
