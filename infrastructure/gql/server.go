package gql

import (
	"context"
	"gqlapi/infrastructure/gql/mutation"
	"gqlapi/infrastructure/gql/query"
	"gqlapi/interface/repository"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type GqlServer struct {
	Port           string
	UserRepo       repository.UserRepository
	CategoriesRepo repository.CategoriesRepository
	ProductRepo    repository.ProductRepository
}

func (g GqlServer) New() {

	schema, err := g.schema()
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			return map[string]interface{}{
				"jwt": r.Header.Get("Authorization"),
			}
		},
	})

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	api := e.Group("/gql")
	api.GET("/playground", echo.WrapHandler(h))
	api.POST("/api", echo.WrapHandler(h))

	e.Logger.Fatal(e.Start(":" + g.Port))
}

func (g GqlServer) schema() (graphql.Schema, error) {
	var schema, err = graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    g.queryDefinition(),
			Mutation: g.mutationDefinition(),
		},
	)
	return schema, err
}

func (g GqlServer) queryDefinition() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				// Auth
				"login": query.Login(g.UserRepo),
				// Categories
				"allcategories":  query.ListCategories(g.CategoriesRepo),
				"categoriesById": query.CategoriesById(g.CategoriesRepo),
				// Products
				"allproducts":  query.ListProducts(g.ProductRepo),
				"productsById": query.ProductById(g.ProductRepo),
			},
		},
	)
}

func (g GqlServer) mutationDefinition() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Mutation",
			Fields: graphql.Fields{
				// Categories
				"createCategories": mutation.CreateCategories(g.CategoriesRepo),
				"deleteCategories": mutation.DeleteCategories(g.CategoriesRepo),
				"updateCategories": mutation.UpdateCategories(g.CategoriesRepo),

				// Products
				"createProducts": mutation.CreateProducts(g.ProductRepo),
				"deleteProducts": mutation.DeleteProducts(g.ProductRepo),
				"updateProducts": mutation.UpdateProducts(g.ProductRepo),
			},
		},
	)
}
