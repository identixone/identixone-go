package utils

import "encoding/json"

func ToMap(e interface{}) (map[string]interface{}, error) {
	var out map[string]interface{}
	data, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
