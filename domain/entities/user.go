package entities

import (
	"encoding/json"
)

type User struct {
	ID       int    `json:"id,omitempty"`
	UserName string `json:"userName,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	JWT      string `json:"jwt,omitempty"`
}

func (u *User) Transcode(data map[string]interface{}) error {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsondata, u); err != nil {
		return err
	}

	return nil
}
