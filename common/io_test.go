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
	Name  string
	Happy bool
	Age   int
}

func TestStruct2Env(t *testing.T) {
	// Struct2Env("MyError", "abc")

	user := &User{Name: "./../Frank", Happy: true, Age: 18}
	Struct2Env("MyUser", user)

	user1 := &User{}
	fPf("%+v\n", *user1)
	user2 := Env2Struct("MyUser", &User{}).(*User)
	fPf("%+v\n", *user2)

	fPln("New Age: ", user2.Age+5)
}

func TestEnv2Struct(t *testing.T) {
	user1 := &User{}
	fPf("%+v\n", *user1)
	user2 := Env2Struct("MyUser", &User{}).(*User)
	fPf("%+v\n", *user2)

	fPln("New Age: ", user2.Age+5)
}
