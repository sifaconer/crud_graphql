package query

import (
	"gqlapi/infrastructure/gql/types"
	"gqlapi/interface/repository"

	"github.com/graphql-go/graphql"
)

func Login(repo repository.UserRepository) *graphql.Field {
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
				return repo.Login(credentials) // call repo implement
			}
			return nil, nil
		},
	}
}
