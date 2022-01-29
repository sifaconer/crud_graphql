package query

import "github.com/graphql-go/graphql"

var QueryDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// Auth
			"login": login(),

			// Categories
			"categories": categories(),

			// Products
			"products": products(),
		},
	},
)
