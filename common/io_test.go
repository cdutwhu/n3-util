package common

import (
	"testing"
)

func TestMustWriteFile(t *testing.T) {
	MustWriteFile("./a/b/c.txt", []byte("hello"))
}

func TestMustAppendFile(t *testing.T) {
	MustAppendFile("./a/b/d.txt", []byte("hello"), true)
}

type User struct {
	Name string
}

func TestStruct2Env(t *testing.T) {
	// Struct2Env("MyError", "abc")

	user := &User{Name: "./../Frank"}
	Struct2Env("MyUser", user)

	user1 := &User{}
	fPf("%+v\n", *user1)
	user2 := Env2Struct("MyUser", &User{}).(*User)
	fPf("%+v\n", *user2)
}

func TestEnv2Struct(t *testing.T) {
	user1 := &User{}
	fPf("%+v\n", *user1)
	user2 := Env2Struct("MyUser", &User{}).(*User)
	fPf("%+v\n", *user2)
}
