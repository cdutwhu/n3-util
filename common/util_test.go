package common

import (
	"testing"
)

func TestColor(t *testing.T) {
	fPln("\033[31mRed")
	fPln("\033[32mGreen")
	fPln("\033[34mBlue")
}

type student struct {
	name  string
	age   int
	score int
}

func (s *student) ShowName(str string) string {
	return str + " " + s.name
}

func (s *student) ShowAge(added int) int {
	return s.age + added
}

func TestTryInvoke(t *testing.T) {
	s := student{name: "HAOHAIDONG", age: 22}
	results, ok, err := TryInvoke(&s, "ShowName", "1")
	if FailOnErr("%v", err); ok {
		Iname, err := InvRst(results, 0)
		FailOnErr("%v", err)
		name := Iname.(string)
		fPln(name)
	}

	// if results, ok := TryInvoke(&s, "ShowAge"); ok {
	// 	age := InvRst(results, 0).(int)
	// 	fPln(age)
	// }
	// TryInvoke(&s, "ShowScore")
}
