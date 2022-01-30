package repository

import "gqlapi/domain/entities"

type UserRepository interface {
	Login(data map[string]interface{}) (entities.Response, error)
}
