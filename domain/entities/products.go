package entities

import "encoding/json"

type Products struct {
	ID          int `json:"id,omitempty"`
	Categories  `json:"categories,omitempty"`
	Name        string   `json:"name,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Description string   `json:"description,omitempty"`
	Stock       *int64   `json:"stock,omitempty"`
}

func (u *Products) Transcode(data map[string]interface{}) error {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsondata, u); err != nil {
		return err
	}

	return nil
}
