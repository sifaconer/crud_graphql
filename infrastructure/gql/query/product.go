package query

import "gqlapi/domain/entities"

func productResolver(id string) (entities.Response, error) {
	var resp entities.Response
	var model entities.Products

	return resp.Ok(model), nil
}
