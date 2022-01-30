package query

import (
	"gqlapi/domain/entities"
	"gqlapi/infrastructure/auth"
	"gqlapi/infrastructure/gql/types"
	"gqlapi/interface/repository"

	"github.com/graphql-go/graphql"
)

func ListCategories(repo repository.CategoriesRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseListCategoriesType,
		Description: "Categories list",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if err := auth.ValidJWT(p.Info.RootValue.(map[string]interface{})["jwt"].(string)); err != nil {
				var resp entities.Response
				return resp.Build(err.Error(), nil).Unauthorized(), err
			}

			return repo.ListCategories()
		},
	}
}

func CategoriesById(repo repository.CategoriesRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseCategoriesType,
		Description: "Categories By ID",
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
				return repo.CategoriesById(id)
			}
			return nil, nil
		},
	}
}
