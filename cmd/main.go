package main

import (
	"gqlapi/infrastructure/gql"
)

func main() {

	s := gql.GqlServer{
		Port: "8182",
	}

	s.New()
}
