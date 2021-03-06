package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sync"
)

const (
	fileEnvName     = "CONF_FILE_PATH"
	defaultFilePath = "conf/dev.json"
)

var values *Values
var mutex = sync.RWMutex{}

// Get returns the config values.
func Get() Values {
	mutex.Lock()
	defer mutex.Unlock()

	if values != nil {
		return *values
	}

	values := &Values{}
	if err := populateDefaults(values); err != nil {
		panic(err)
	}
	if err := overrideWithFile(values, getFilePath()); err != nil {
		panic(err)
	}
	return *values
}

func populateDefaults(target interface{}) error {
	if !isPointer2Struct(target) {
		return errors.New("provided argument is not pointer to struct")
	}

	structValue := reflect.ValueOf(target).Elem().Interface()
	defaultsMap, err := createDefaultsMap4Struct(structValue)
	if err != nil {
		return err
	}

	defaultsJSON, err := json.Marshal(defaultsMap)
	if err != nil {
		return err
	}

	return json.Unmarshal(defaultsJSON, target)
}

func overrideWithFile(target interface{}, filePath string) error {
	fileContents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileContents, target)
}

func createDefaultsMap4Struct(input interface{}) (interface{}, error) {
	inputValue := reflect.ValueOf(input)
	inputType := reflect.TypeOf(input)

	defaultsMap := map[string]interface{}{}

	for i := 0; i < inputValue.NumField(); i++ {
		fieldValue := inputValue.Field(i)
		fieldType := inputType.Field(i)

		keyName, present := fieldType.Tag.Lookup("json")
		if !present || keyName == "" {
			return nil, errors.New("struct field does not have 'json' tag")
		}
		defaultValue, present := fieldType.Tag.Lookup("default")
		if present {
			parsed, err := parseString2Interface(fieldValue.Kind(), defaultValue)
			if err != nil {
				return nil, fmt.Errorf("default value is invalid for field '%s'", keyName)
			}
			defaultsMap[keyName] = parsed
		}

		if fieldValue.Kind() == reflect.Struct {
			internalValue, err := createDefaultsMap4Struct(fieldValue.Interface())
			if err != nil {
				return nil, fmt.Errorf("default value is invalid for internal fields of struct: '%s'", keyName)
			}
			defaultsMap[keyName] = internalValue
		}
	}
	return defaultsMap, nil
}

func getFilePath() string {
	filePath := os.Getenv(fileEnvName)
	if filePath == "" {
		return defaultFilePath
	}
	return filePath
}
