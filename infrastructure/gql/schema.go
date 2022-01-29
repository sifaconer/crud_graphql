package gql

import (
	"gqlapi/infrastructure/gql/mutation"
	"gqlapi/infrastructure/gql/query"

	"github.com/graphql-go/graphql"
)

func schema() (graphql.Schema, error) {
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    query.QueryDefinition,
			Mutation: mutation.MutationDefinition,
		},
	)
	return schema, err
}
