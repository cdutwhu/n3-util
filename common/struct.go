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

// Struct2Map : each field name MUST be Exportable
func Struct2Map(s interface{}) (map[string]interface{}, error) {
	stVal := reflect.ValueOf(s)
	if stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct {
		return nil, eg.PARAM_INVALID_STRUCT_PTR
	}
	ret := make(map[string]interface{})
	stValElem := stVal.Elem()
	valTyp := stValElem.Type()
	for i := 0; i < stValElem.NumField(); i++ {
		if name, field := valTyp.Field(i).Name, stValElem.Field(i); field.CanInterface() {
			ret[name] = field.Interface()

			// --------------- //
			if field.Type().Kind() == reflect.Func {
				fPln("func variable: " + name)
			}
			// --------------- //
		}
	}
	return ret, nil
}

// StructFields :
func StructFields(s interface{}) ([]string, error) {
	stVal := reflect.ValueOf(s)
	if stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct {
		return nil, eg.PARAM_INVALID_STRUCT_PTR
	}
	m, err := Struct2Map(s)
	FailOnErr("%v", err)
	IKeys, err := MapKeys(m)
	FailOnErr("%v", err)
	return IKeys.([]string), nil
}
