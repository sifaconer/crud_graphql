package database

import (
	"context"
	"gqlapi/domain/entities"
	"gqlapi/infrastructure/database/ent"
	"gqlapi/infrastructure/database/ent/categories"
)

type CategoriesImpl struct {
	DB  *ent.Client
	Ctx context.Context
}

func (c CategoriesImpl) ListCategories() (entities.Response, error) {
	var resp entities.Response
	var model []entities.Categories

	allCategories, err := c.DB.Categories.
		Query().
		All(c.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	for _, v := range allCategories {
		model = append(model, entities.Categories{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return resp.Build("OK", model).Ok(), nil
}

func (c CategoriesImpl) CategoriesById(id int) (entities.Response, error) {
	var resp entities.Response
	var model entities.Categories

	idCategories, err := c.DB.Categories.
		Query().Where(categories.ID(id)).
		Only(c.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return resp.Build("OK", nil).Ok(), nil
		}
		return resp.Build(err.Error(), nil).Bad(), err
	}

	model = entities.Categories{
		ID:          idCategories.ID,
		Name:        idCategories.Name,
		Description: idCategories.Description,
	}

	return resp.Build("OK", model).Ok(), nil
}

func (c CategoriesImpl) CreateCategoriesResolver(params map[string]interface{}) (entities.Response, error) {
	var resp entities.Response

	var model entities.Categories
	if err := model.Transcode(params); err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	newCategories, err := c.DB.Categories.
		Create().
		SetName(model.Name).
		SetDescription(model.Description).Save(c.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	return resp.Build("Categoria Creada", entities.Categories{
		ID:          newCategories.ID,
		Name:        newCategories.Name,
		Description: newCategories.Description,
	}).Ok(), nil
}

func (c CategoriesImpl) DeleteCategoriesResolver(id int) (entities.Response, error) {
	var resp entities.Response

	if err := c.DB.Categories.
		DeleteOneID(id).
		Exec(c.Ctx); err != nil {
		if ent.IsNotFound(err) {
			return resp.Build("Categoria Borrada", nil).Ok(), nil
		}
		return resp.Build(err.Error(), nil).Bad(), err
	}

	return resp.Build("Categoria Borrada", nil).Ok(), nil
}

func (c CategoriesImpl) UpdateCategoriesResolver(params map[string]interface{}) (entities.Response, error) {
	var resp entities.Response

	var model entities.Categories
	if err := model.Transcode(params); err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	updateCategories, err := c.queryBuilder(
		c.DB.Categories.UpdateOneID(model.ID),
		model).
		Save(c.Ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return resp.Build("Categoria Actualizada", nil).Ok(), nil
		}
		return resp.Build(err.Error(), nil).Bad(), err
	}

	return resp.Build("Categoria Actualizada", entities.Categories{
		ID:          updateCategories.ID,
		Name:        updateCategories.Name,
		Description: updateCategories.Description,
	}).Ok(), nil
}

func (c CategoriesImpl) queryBuilder(categoriesUpdate *ent.CategoriesUpdateOne, model entities.Categories) *ent.CategoriesUpdateOne {

	if len(model.Name) != 0 {
		categoriesUpdate.SetName(model.Name)
	}
	if len(model.Description) != 0 {
		categoriesUpdate.SetDescription(model.Description)
	}

	return categoriesUpdate
}
