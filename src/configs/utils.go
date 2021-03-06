package configs

import (
	"encoding/json"
	"reflect"
)

func isPointer2Struct(input interface{}) bool {
	value := reflect.ValueOf(input)
	if value.Kind() != reflect.Ptr {
		return false
	}
	return value.Elem().Kind() == reflect.Struct
}

func parseString2Interface(kind reflect.Kind, value string) (interface{}, error) {
	if kind == reflect.String {
		return value, nil
	}
	var converted interface{}
	if err := json.Unmarshal([]byte(value), &converted); err != nil {
		return nil, err
	}
	return converted, nil
}
