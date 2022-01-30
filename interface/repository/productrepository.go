package repository

import "gqlapi/domain/entities"

type ProductRepository interface {
	ListProduct() (entities.Response, error)
	ProductById(id int) (entities.Response, error)
	CreateProductsResolver(params map[string]interface{}) (entities.Response, error)
	DeleteProductsResolver(id int) (entities.Response, error)
	UpdateProductsResolver(params map[string]interface{}) (entities.Response, error)
}
