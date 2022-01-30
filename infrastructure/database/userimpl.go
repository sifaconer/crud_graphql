package database

import (
	"context"
	"errors"
	"gqlapi/domain/entities"
	"gqlapi/infrastructure/auth"
	"gqlapi/infrastructure/database/ent"
	"gqlapi/infrastructure/database/ent/user"
	"gqlapi/utils"
	"strings"
)

type UserImpl struct {
	DB  *ent.Client
	Ctx context.Context
}

func (u UserImpl) Login(data map[string]interface{}) (entities.Response, error) {
	var resp entities.Response
	var model entities.User
	err := model.Transcode(data)
	if err != nil {
		return resp.Build(err.Error(), nil).Bad(), err
	}

	user, err := u.DB.User.
		Query().
		Where(
			user.And(
				user.UserName(model.UserName),
				user.Password(utils.HashPassword(strings.TrimSpace(model.Password))),
			),
		).
		Only(u.Ctx)
	if err != nil {
		return resp.Build(err.Error(), nil).
			Unauthorized(), errors.New("credentials incorrect")
	}

	if user == nil {
		return resp.Build("Credentials incorrect", nil).
			Unauthorized(), errors.New("credentials incorrect")
	}

	if !utils.CheckPasswordHash(model.Password, user.Password) {
		return resp.Build("Credentials incorrect", nil).
			Unauthorized(), errors.New("credentials incorrect")
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return resp.Build(err.Error(), nil).
			IntervalServerError(), errors.New("jwt error")
	}

	return resp.Build("OK", entities.User{
		ID:       user.ID,
		UserName: user.UserName,
		Name:     user.Name,
		JWT:      token,
	}).Ok(), nil
}
