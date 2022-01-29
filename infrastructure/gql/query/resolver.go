package query

import (
	"gqlapi/infrastructure/gql/types"

	"github.com/graphql-go/graphql"
)

func login() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseUserType,
		Description: "Login user",
		Args: graphql.FieldConfigArgument{
			"credentials": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.UserInputType),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if credentials, ok := p.Args["credentials"].(map[string]interface{}); ok {
				return loginResolver(credentials)
			}
			return nil, nil
		},
	}
}

func categories() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseCategoriesType,
		Description: "Categories list",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, ok := p.Args["id"].(string); ok {
				return categoriesResolver(id)
			}
			return nil, nil
		},
	}
}

func products() *graphql.Field {
	return &graphql.Field{
		Type:        types.ResponseProductType,
		Description: "Products list",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if id, ok := p.Args["id"].(string); ok {
				return productResolver(id)
			}
			return nil, nil
		},
	}
}
