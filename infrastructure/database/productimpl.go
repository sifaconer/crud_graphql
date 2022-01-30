package database

import (
	"context"
	"gqlapi/domain/entities"
	"gqlapi/infrastructure/database/ent"
	"gqlapi/infrastructure/database/ent/product"
)

type ProductImpl struct {
	DB  *ent.Client
	Ctx context.Context
}

func (p ProductImpl) ListProduct() (entities.Response, error) {
	var resp entities.Response
	var model []entities.Products

	allProduct, err := p.DB.Product.
		Query().
		All(p.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	for _, v := range allProduct {
		pc, err := v.QueryCategories().Only(p.Ctx)
		if err != nil {
			return resp.Build(err.Error(), nil).Bad(), err
		}
		model = append(model, entities.Products{
			ID: v.ID,
			Categories: entities.Categories{
				ID:          pc.ID,
				Name:        pc.Name,
				Description: pc.Description,
			},
			Price:       &v.Price,
			Stock:       &v.Stock,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return resp.Build("OK", model).Ok(), nil
}

func (p ProductImpl) ProductById(id int) (entities.Response, error) {
	var resp entities.Response
	var model entities.Products

	idProduct, err := p.DB.Product.
		Query().Where(product.ID(id)).
		Only(p.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return resp.Build("OK", nil).Ok(), nil
		}
		return resp.Build(err.Error(), nil).Bad(), nil
	}

	pc, err := idProduct.QueryCategories().Only(p.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	model = entities.Products{
		ID: idProduct.ID,
		Categories: entities.Categories{
			ID:          pc.ID,
			Name:        pc.Name,
			Description: pc.Description,
		},
		Price:       &idProduct.Price,
		Stock:       &idProduct.Stock,
		Name:        idProduct.Name,
		Description: idProduct.Description,
	}

	return resp.Build("OK", model).Ok(), nil
}

func (p ProductImpl) CreateProductsResolver(params map[string]interface{}) (entities.Response, error) {
	var resp entities.Response

	var model entities.Products
	if err := model.Transcode(params); err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	newProduct, err := p.DB.Product.
		Create().
		SetCategoriesID(model.Categories.ID).
		SetPrice(*model.Price).
		SetStock(*model.Stock).
		SetName(model.Name).
		SetDescription(model.Description).Save(p.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	pc, err := newProduct.QueryCategories().First(p.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}
	return resp.Build("Product Creado", entities.Products{
		ID: newProduct.ID,
		Categories: entities.Categories{
			ID:          pc.ID,
			Name:        pc.Name,
			Description: pc.Description,
		},
		Price:       &newProduct.Price,
		Stock:       &newProduct.Stock,
		Name:        newProduct.Name,
		Description: newProduct.Description,
	}).Ok(), nil
}

func (p ProductImpl) DeleteProductsResolver(id int) (entities.Response, error) {
	var resp entities.Response

	if err := p.DB.Product.
		DeleteOneID(id).
		Exec(p.Ctx); err != nil {
		if ent.IsNotFound(err) {
			return resp.Build("Product Borrado", nil).Ok(), nil
		}
		return resp.Build(err.Error(), nil).Bad(), err
	}

	return resp.Build("Product Borrado", nil).Ok(), nil
}

func (p ProductImpl) UpdateProductsResolver(params map[string]interface{}) (entities.Response, error) {
	var resp entities.Response

	var model entities.Products
	if err := model.Transcode(params); err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	updateProduct, err := p.queryBuilder(
		p.DB.Product.UpdateOneID(model.ID),
		model).
		Save(p.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return resp.Build("Product Actualizado", nil).Ok(), nil
		}
		return resp.Build(err.Error(), nil).Bad(), err
	}

	pc, err := updateProduct.QueryCategories().Only(p.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	return resp.Build("Product Actualizado", entities.Products{
		ID: updateProduct.ID,
		Categories: entities.Categories{
			ID:          pc.ID,
			Name:        pc.Name,
			Description: pc.Description,
		},
		Price:       &updateProduct.Price,
		Stock:       &updateProduct.Stock,
		Name:        updateProduct.Name,
		Description: updateProduct.Description,
	}), nil
}

func (p ProductImpl) queryBuilder(productUpdate *ent.ProductUpdateOne, model entities.Products) *ent.ProductUpdateOne {

	if model.Categories.ID != 0 {
		productUpdate.SetCategoriesID(model.Categories.ID)
	}

	if model.Price != nil {
		productUpdate.SetPrice(*model.Price)
	}

	if model.Stock != nil {
		productUpdate.SetStock(*model.Stock)
	}

	if len(model.Name) != 0 {
		productUpdate.SetName(model.Name)
	}

	if len(model.Description) != 0 {
		productUpdate.SetDescription(model.Description)
	}

	return productUpdate
}
