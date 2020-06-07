package common

import "testing"

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
	Iuser2, err := Env2Struct("MyUser", &User{})
	FailOnErr("%v", err)
	user2 := Iuser2.(*User)
	fPf("%+v\n", *user2)

	fPln("New Age: ", user2.Age+5)
}

func TestEnv2Struct(t *testing.T) {
	user1 := &User{}
	fPf("%+v\n", *user1)
	Iuser2, err := Env2Struct("MyUser", &User{})
	FailOnErr("%v", err)
	user2 := Iuser2.(*User)
	fPf("%+v\n", *user2)

	fPln("New Age: ", user2.Age+5)
}

func TestStruct2Map(t *testing.T) {
	s := struct {
		A string
		B int
		C bool
	}{A: "aa", B: 22, C: false}
	m, err := Struct2Map(s)
	FailOnErr("%v", err)
	fPln(m)

	// ERROR !!!
	// ks, vs := MapKVs(m)
	// fPln(ks)
	// fPln(vs)

	fPln(" ------------------- ")

	s1 := struct {
		A1 string
		B1 string
		C1 string
	}{A1: "aa", B1: "22", C1: "false"}
	m, err = Struct2Map(s1)
	FailOnErr("%v", err)
	fPln(m)

	ks, vs, err := MapKVs(m)
	FailOnErr("%v", err)
	fPln(ks, vs)
}

func TestStructFields(t *testing.T) {
	s := struct {
		A string
		B int
		C bool
	}{A: "aa", B: 22, C: false}
	flds := StructFields(s)
	fPln(flds)
}
