package mutation

import (
	"gqlapi/domain/entities"
	"gqlapi/infrastructure/auth"
	"gqlapi/infrastructure/gql/types"
	"gqlapi/interface/repository"

	"github.com/graphql-go/graphql"
)

// CATEGORIES
func CreateCategories(repo repository.CategoriesRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseCategoriesType,
		Description: "Create categories",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.CategoriesCreateInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if err := auth.ValidJWT(p.Info.RootValue.(map[string]interface{})["jwt"].(string)); err != nil {
				var resp entities.Response
				return resp.Build(err.Error(), nil).Unauthorized(), err
			}

			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return repo.CreateCategoriesResolver(params)
			}
			return nil, nil
		},
	}
}
func DeleteCategories(repo repository.CategoriesRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseType,
		Description: "Delete categories",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if err := auth.ValidJWT(p.Info.RootValue.(map[string]interface{})["jwt"].(string)); err != nil {
				var resp entities.Response
				return resp.Build(err.Error(), nil).Unauthorized(), err
			}

			if id, ok := p.Args["id"].(int); ok {
				return repo.DeleteCategoriesResolver(id)
			}
			return nil, nil
		},
	}
}
func UpdateCategories(repo repository.CategoriesRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseCategoriesType,
		Description: "Update categories",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.CategoriesInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if err := auth.ValidJWT(p.Info.RootValue.(map[string]interface{})["jwt"].(string)); err != nil {
				var resp entities.Response
				return resp.Build(err.Error(), nil).Unauthorized(), err
			}

			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return repo.UpdateCategoriesResolver(params)
			}
			return nil, nil
		},
	}
}
