package mutation

import (
	"gqlapi/domain/entities"
	"gqlapi/infrastructure/auth"
	"gqlapi/infrastructure/gql/types"
	"gqlapi/interface/repository"

	"github.com/graphql-go/graphql"
)

// PRODUCT
func CreateProducts(repo repository.ProductRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseProductType,
		Description: "Create product",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.ProductCreateInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if err := auth.ValidJWT(p.Info.RootValue.(map[string]interface{})["jwt"].(string)); err != nil {
				var resp entities.Response
				return resp.Build(err.Error(), nil).Unauthorized(), err
			}

			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return repo.CreateProductsResolver(params)
			}
			return nil, nil
		},
	}
}

func DeleteProducts(repo repository.ProductRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseType,
		Description: "Delete product",
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
				return repo.DeleteProductsResolver(id)
			}
			return nil, nil
		},
	}
}

func UpdateProducts(repo repository.ProductRepository) *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseProductType,
		Description: "Update product",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.ProductInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if err := auth.ValidJWT(p.Info.RootValue.(map[string]interface{})["jwt"].(string)); err != nil {
				var resp entities.Response
				return resp.Build(err.Error(), nil).Unauthorized(), err
			}

			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return repo.UpdateProductsResolver(params)
			}
			return nil, nil
		},
	}
}
