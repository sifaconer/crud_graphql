package mutation

import "github.com/graphql-go/graphql"

var MutationDefinition = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// Categories
			"createCategories": createCategories(),
			"deleteCategories": deleteCategories(),
			"updateCategories": updateCategories(),

			// Products
			"createProducts": createProducts(),
			"deleteProducts": deleteProducts(),
			"updateProducts": updateProducts(),
		},
	},
)
