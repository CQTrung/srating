package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSONB []interface{}

type JSONA interface{}

func (a JSONB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &a)
}

func ValidateCodeUniqueness(data JSONB) bool {
	encounteredCodes := make(map[string]bool)
	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			return false
		}
		code, ok := itemMap["code"].(string)
		if !ok {
			return false
		}
		if encounteredCodes[code] {
			return false // Code is not unique
		}
		encounteredCodes[code] = true
	}
	return true // All codes are unique
}
