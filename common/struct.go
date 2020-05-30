package common

import (
	"encoding/json"
	"os"
	"reflect"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// Struct2Env :
func Struct2Env(key string, s interface{}) {
	FailOnErrWhen(reflect.ValueOf(s).Kind() != reflect.Ptr, "%v", eg.PARAM_INVALID_PTR)
	FailOnErrWhen(reflect.ValueOf(s).Elem().Kind() != reflect.Struct, "%v", eg.PARAM_INVALID_STRUCT)
	bytes, err := json.Marshal(s)
	FailOnErr("%v", err)
	FailOnErr("%v", os.Setenv(key, string(bytes)))
}

// Env2Struct :
func Env2Struct(key string, s interface{}) interface{} {
	FailOnErrWhen(reflect.ValueOf(s).Kind() != reflect.Ptr, "%v", eg.PARAM_INVALID_PTR)
	FailOnErrWhen(reflect.ValueOf(s).Elem().Kind() != reflect.Struct, "%v", eg.PARAM_INVALID_STRUCT)
	jsonstr := os.Getenv(key)
	FailOnErrWhen(!IsJSON(jsonstr), "%v", eg.JSON_INVALID)
	FailOnErr("%v", json.Unmarshal([]byte(jsonstr), s))
	return s
}

// Struct2Map : each field name MUST be exportable
func Struct2Map(s interface{}) map[string]interface{} {
	FailOnErrWhen(reflect.ValueOf(s).Kind() != reflect.Struct, "%v", eg.PARAM_INVALID_STRUCT)
	ret := make(map[string]interface{})
	v := reflect.ValueOf(s)
	val := reflect.Indirect(v)
	for i := 0; i < v.NumField(); i++ {
		name := val.Type().Field(i).Name
		field := v.Field(i).Interface()
		ret[name] = field
	}
	return ret
}

// StructFields :
func StructFields(s interface{}) []string {
	return MapKeys(Struct2Map(s)).([]string)
}
