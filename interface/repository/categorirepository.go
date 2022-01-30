package repository

import "gqlapi/domain/entities"

type CategoriesRepository interface {
	ListCategories() (entities.Response, error)
	CategoriesById(id int) (entities.Response, error)
	CreateCategoriesResolver(params map[string]interface{}) (entities.Response, error)
	DeleteCategoriesResolver(id int) (entities.Response, error)
	UpdateCategoriesResolver(params map[string]interface{}) (entities.Response, error)
}
