package common

import (
	"github.com/golang-api-user/user/models"
	"encoding/json"
)

// CreateFromMap function for convert map to struct
func CreateFromMap(m map[string]interface{}) (*models.User, error) {
	data, _ := json.Marshal(m)
	var result = new(models.User)
	err := json.Unmarshal(data, &result)
	return result, err
}
