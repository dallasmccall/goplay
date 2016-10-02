package Collectors

import (
	"encoding/json"
)

func ToJSON(data interface{}) (string, error) {
	bytes, err := json.MarshalIndent(data, "", " ")
	if nil != err {
		return "", err
	}
	return string(bytes), nil
}
