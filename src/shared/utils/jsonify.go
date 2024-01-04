package utils

import (
	"encoding/json"
	"htmx/src/shared/utils/logger"
)

func JsonifyPretty(data interface{}) string {
	dataJson, err := json.MarshalIndent(data, "", "  ")

	if err != nil {
		msg := "Error while marshalling user: " + err.Error()
		logger.Default.Error(msg)
		return "{}"
	}

	return string(dataJson)
}

func Jsonify(data interface{}) string {
	dataJson, err := json.Marshal(data)

	if err != nil {
		msg := "Error while marshalling user: " + err.Error()
		logger.Default.Error(msg)
		return "{}"
	}

	return string(dataJson)
}
