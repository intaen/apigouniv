package util

import "encoding/json"

// ConvertResponse is a function to change struct to map interface like response format
func ConvertResponse(data interface{}) (map[string]interface{}, string) {
	var newData map[string]interface{}
	dt, _ := json.Marshal(data)
	json.Unmarshal(dt, &newData)

	return newData, string(dt)
}
