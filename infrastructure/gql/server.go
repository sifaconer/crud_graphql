package gql

import (
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

type GqlServer struct {
	Port string
}

func (g *GqlServer) New() {

	schema, err := schema()
	if err != nil {
		log.Fatalln(err)
	}

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	http.Handle("/api/gql/playground", h)
	if err := http.ListenAndServe(":"+g.Port, nil); err != nil {
		log.Fatalln(err)
	}
}
