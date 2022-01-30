package main

import (
	"context"
	"gqlapi/infrastructure/database"
	"gqlapi/infrastructure/gql"
	"log"
)

func main() {

	ctx := context.Background()
	uriDB := "host=127.0.0.1 port=5432 user=root dbname=postgres password=root sslmode=disable"
	/* if uriDB = os.Getenv("URI_DATABASE"); len(uriDB) == 0 {
		log.Fatalf("ENV-ERROR: Environment variable [URI_DATABASE] is empty")
	} */

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
