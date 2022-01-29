package entities

import "encoding/json"

type Categories struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (u *Categories) Transcode(data map[string]interface{}) error {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsondata, u); err != nil {
		return err
	}

	return nil
}
