package common

import (
	"encoding/json"
	"os"
	"reflect"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// Struct2Env :
func Struct2Env(key string, s interface{}) error {
	stVal := reflect.ValueOf(s)
	if stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct {
		return eg.PARAM_INVALID_STRUCT_PTR
	}
	bytes, err := json.Marshal(s)
	FailOnErr("%v", err)
	FailOnErr("%v", os.Setenv(key, string(bytes)))
	return nil
}

// Env2Struct :
func Env2Struct(key string, s interface{}) (interface{}, error) {
	stVal := reflect.ValueOf(s)
	if stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct {
		return nil, eg.PARAM_INVALID_STRUCT_PTR
	}
	jsonstr := os.Getenv(key)
	FailOnErrWhen(!IsJSON(jsonstr), "%v", eg.JSON_INVALID)
	FailOnErr("%v", json.Unmarshal([]byte(jsonstr), s))
	return s, nil
}

// Struct2Map : each field name MUST be exportable
func Struct2Map(s interface{}) (map[string]interface{}, error) {
	stVal := reflect.ValueOf(s)
	if stVal.Kind() != reflect.Struct {
		return nil, eg.PARAM_INVALID_STRUCT
	}
	ret := make(map[string]interface{})
	val := reflect.Indirect(stVal)
	for i := 0; i < stVal.NumField(); i++ {
		name := val.Type().Field(i).Name
		field := stVal.Field(i).Interface()
		ret[name] = field
	}
	return ret, nil
}

// StructFields :
func StructFields(s interface{}) []string {
	m, err := Struct2Map(s)
	FailOnErr("%v", err)
	keys, err := MapKeys(m)
	FailOnErr("%v", err)
	return keys.([]string)
}
