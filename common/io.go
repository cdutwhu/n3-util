package common

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// MustWriteFile :
func MustWriteFile(filename string, data []byte) {
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
	FailOnErr("%v", ioutil.WriteFile(filename, data, 0666))
}

// MustAppendFile :
func MustAppendFile(filename string, data []byte, newline bool) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		MustWriteFile(filename, []byte(""))
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	FailOnErr("%v", err)
	defer file.Close()

	if newline {
		data = append(append([]byte{}, '\n'), data...)
	}
	_, err = file.Write(data)
	FailOnErr("%v", err)
}

// Struct2Env :
func Struct2Env(key string, s interface{}) {
	FailOnErrWhen(reflect.ValueOf(s).Kind() != reflect.Ptr, "%v", eg.PARAM_INVALID_PTR)
	FailOnErrWhen(reflect.ValueOf(s).Elem().Kind() != reflect.Struct, "%v", eg.PARAM_INVALID_STRUCT)
	bytes, err := json.Marshal(s)
	FailOnErr("%v", err)
	FailOnErr("%v", os.Setenv(key, string(bytes)))
}

// Env2Struct :
func Env2Struct(key string, s interface{}) {
	FailOnErrWhen(reflect.ValueOf(s).Kind() != reflect.Ptr, "%v", eg.PARAM_INVALID_PTR)
	FailOnErrWhen(reflect.ValueOf(s).Elem().Kind() != reflect.Struct, "%v", eg.PARAM_INVALID_STRUCT)
	jsonstr := os.Getenv(key)
	FailOnErrWhen(!IsJSON(jsonstr), "%v", eg.JSON_INVALID)
	FailOnErr("%v", json.Unmarshal([]byte(jsonstr), s))
}
