package query

import "gqlapi/domain/entities"

func categoriesResolver(id string) (entities.Response, error) {
	var resp entities.Response
	var model entities.Categories

	return resp.Ok(model), nil
}
