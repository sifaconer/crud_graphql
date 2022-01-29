package mutation

import (
	"gqlapi/infrastructure/gql/types"

	"github.com/graphql-go/graphql"
)

// CATEGORIES
func createCategories() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseCategoriesType,
		Description: "Create categories",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.CategoriesInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return createCategoriesResolver(params)
			}
			return nil, nil
		},
	}
}
func deleteCategories() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseType,
		Description: "Delete categories",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, ok := p.Args["id"].(string); ok {
				return deleteCategoriesResolver(id)
			}
			return nil, nil
		},
	}
}
func updateCategories() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseCategoriesType,
		Description: "Update categories",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.CategoriesInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return updateCategoriesResolver(params)
			}
			return nil, nil
		},
	}
}

// PRODUCT
func createProducts() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseProductType,
		Description: "Create product",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.ProductInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return createProductsResolver(params)
			}
			return nil, nil
		},
	}
}
func deleteProducts() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseType,
		Description: "Delete product",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, ok := p.Args["id"].(string); ok {
				return deleteProductsResolver(id)
			}
			return nil, nil
		},
	}
}
func updateProducts() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseProductType,
		Description: "Update product",
		Args: graphql.FieldConfigArgument{
			"params": &graphql.ArgumentConfig{
				Type: types.ProductInputType,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if params, ok := p.Args["params"].(map[string]interface{}); ok {
				return updateProductsResolver(params)
			}
			return nil, nil
		},
	}
}
