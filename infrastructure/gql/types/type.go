package types

import "github.com/graphql-go/graphql"

// RESPONSE
var ResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResponseType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"statusCode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// USER
var ResponseUserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResponseUserType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"statusCode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: userType,
			},
		},
	},
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userName": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"jwt": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var UserInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"userName": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// CATEGORIES
var ResponseCategoriesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResponseCategoriesType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"statusCode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: categoriesType,
			},
		},
	},
)
var ResponseListCategoriesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResponseListCategoriesType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"statusCode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: graphql.NewList(categoriesType),
			},
		},
	},
)

var categoriesType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Categories",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
var CategoriesInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CategoriesInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)
var CategoriesCreateInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "CategoriesCreateInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

// PRODUCTS
var ResponseProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResponseProductType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"statusCode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: productType,
			},
		},
	},
)
var ResponseListProductType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ResponseListProductType",
		Fields: graphql.Fields{
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"statusCode": &graphql.Field{
				Type: graphql.Int,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"data": &graphql.Field{
				Type: graphql.NewList(productType),
			},
		},
	},
)

var productType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Product",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"categories": &graphql.Field{
				Type: categoriesType,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"price": &graphql.Field{
				Type: graphql.Float,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"stock": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
var ProductCreateInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ProductCreateInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"categories": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(CategoriesInputType),
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Float),
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"stock": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
		},
	},
)
var ProductInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "ProductInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"id": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"categories": &graphql.InputObjectFieldConfig{
				Type: CategoriesInputType,
			},
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"price": &graphql.InputObjectFieldConfig{
				Type: graphql.Float,
			},
			"description": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"stock": &graphql.InputObjectFieldConfig{
				Type: graphql.Int,
			},
		},
	},
)
