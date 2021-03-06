package main

import (
	"context"
	"gqlapi/infrastructure/database"
	"gqlapi/infrastructure/gql"
	"log"
	"os"
)

func main() {

	ctx := context.Background()
	uriDB := ""
	if uriDB = os.Getenv("URI_DATABASE"); len(uriDB) == 0 {
		log.Fatalf("ENV-ERROR: Environment variable [URI_DATABASE] is empty")
	}

	db := database.Database{
		UriDB: uriDB,
		Ctx:   ctx,
	}
	client, err := db.New()
	if err != nil {
		log.Fatalf("ERROR IN DATABASE: %v", err)
	}
	defer client.Close()

	s := gql.GqlServer{
		Port: "8182",
		UserRepo: database.UserImpl{
			DB:  client,
			Ctx: ctx,
		},
		CategoriesRepo: database.CategoriesImpl{
			DB:  client,
			Ctx: ctx,
		},
		ProductRepo: database.ProductImpl{
			DB:  client,
			Ctx: ctx,
		},
	}

	s.New()
}
