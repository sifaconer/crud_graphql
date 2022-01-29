package query

import "gqlapi/domain/entities"

func loginResolver(data map[string]interface{}) (entities.Response, error) {
	var resp entities.Response
	var user entities.User
	err := user.Transcode(data)
	if err != nil {
		resp.Message = err.Error()
		return resp.Bad(), err
	}
	return resp.Ok(user), nil
}
