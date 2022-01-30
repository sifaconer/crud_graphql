package database

import (
	"context"
	"gqlapi/infrastructure/database/ent"
	"gqlapi/utils"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	UriDB string
	Ctx   context.Context
}

func (d *Database) New() (*ent.Client, error) {
	client, err := ent.Open("postgres", d.UriDB)
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(
		d.Ctx,
	); err != nil {
		return nil, err
	}

	_, err = client.User.
		Create().
		SetName("Usuario 1").
		SetUserName("user1").
		SetPassword(utils.HashPassword("$user1*")).Save(d.Ctx)
	if err != nil {
		log.Printf("WARNING: error creando usuario por defecto: %s\n", err.Error())
	}

	return client, nil
}
