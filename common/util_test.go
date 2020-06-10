package common

import (
	"testing"
)

func TestColor(t *testing.T) {
	fPln("\033[31mRed")
	fPln("\033[32mGreen")
	fPln("\033[34mBlue")
}

// ------------------------- //

// Iperson :
type Iperson interface {
	ShowName(s string) string
}

type Person struct {
	Name string
	Age  int
	Fn   func()
}

func (p *Person) ShowName(str string) string {
	return str + " P " + p.Name
}

func (p *Person) ShowAge(added int) int {
	return p.Age + added
}

type Student struct {
	Person
	score int
	MW    map[string]map[string][]interface{}
}

func (s *Student) ShowName(str string) string {
	return str + " S " + s.Name
}

func (s *Student) ShowScore(str string) {
	fPt(str + "   ")
	fPln(s.score)
}

func (s *Student) AddScore(added int) {
	fPln(s.score + added)
}

// Show :
func Show(ip Iperson) {
	fPln(ip.ShowName("hello"))
}

func TestTryInvoke(t *testing.T) {
	s := &Student{
		Person: Person{
			Name: "HAOHAIDONG",
			Age:  22,
		},
		score: 100,
		MW: map[string]map[string][]interface{}{
			"ShowScore": {
				"$@":       {"$1"},
				"ShowName": {"$1"},
			},
			// "AddScore": {
			// 	"$@":       {1000},
			// 	"ShowName": {500},
			// },
		},
	}

	Show(s)

	results, ok, err := TryInvokeWithMW(s, "ShowName", "Great")
	if FailOnErr("%v", err); ok {
		Iname, err := InvRst(results, 0)
		FailOnErr("%v", err)
		name := Iname.(string)
		fPln(name)
	}
}
