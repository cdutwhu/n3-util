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
	ShowName(s1, s2 string) string
}

type Person struct {
	Name string
	Age  int
	Fn   func()
}

func (p *Person) ShowName(str1, str2 string) string {
	return str1 + " P " + str2 + " P " + p.Name
}

func (p *Person) ShowAge(added int) int {
	return p.Age + added
}

type Student struct {
	Person
	score int
	MW    map[string]map[string][]interface{}
}

// func (s *Student) ShowName(str1, str2 string) string {
// 	return str1 + " S " + str2 + " S " + s.Name
// }

func (s *Student) ShowScore(str string) {
	fPt(str + "   ")
	fPln(s.score)
}

func (s *Student) AddScore(added int) {
	fPln(s.score + added)
}

// Show :
func Show(ip Iperson) {
	fPln(ip.ShowName("hello", "world"))
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
				"*":        {"$1"},
				"ShowName": {"$@"},
			},
			// "AddScore": {
			// 	"$@":       {1000},
			// 	"ShowName": {500},
			// },
		},
	}

	Show(s)

	fPln(MustInvokeWithMW(s, "ShowName", "Great", "haohaidong"))

	// results, ok, err := TryInvokeWithMW(s, "ShowName", "Great")
	// if FailOnErr("%v", err); ok {
	// 	Iname, err := InvRst(results, 0)
	// 	FailOnErr("%v", err)
	// 	name := Iname.(string)
	// 	fPln(name)
	// }
}
