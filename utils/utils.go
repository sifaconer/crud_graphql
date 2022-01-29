package utils

import "encoding/json"

// Transcode : params=data input graphql, model=struct model to load data
func Transcode(params map[string]interface{}, model interface{}) error {
	jsondata, err := json.Marshal(params)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsondata, model); err != nil {
		return err
	}

	return nil
}
