package common

import "testing"

func TestRmTailFromLast(t *testing.T) {
	fPln(RmTailFromLast("AB.CD.EF", "."))
	fPln(RmTailFromLast("AB.CD.EF", "#"))
}

func TestRmTailFromLastN(t *testing.T) {
	fPln(RmTailFromLastN("AB.CD.EF", ".", 2))
	fPln(RmTailFromLastN("AB.CD.EF", "#", 2))
}

func TestRmTailFromFirst(t *testing.T) {
	fPln(RmTailFromFirstAny(`Activity>RefId="C27E1FCF-C163-485F-BEF0-F36F18A0493A" lang="en"`, " ", ">"))
}

func TestRmHeadToLast(t *testing.T) {
	fPln(RmHeadToLast("##AB##CD##F", "D"))
}

func TestRmHeadToFirst(t *testing.T) {
	fPln(RmHeadToFirst("777##AB##CD##F", "##"))
}

func TestStrReplByPos(t *testing.T) {
	s, err := ReplByPosGrp("0123456789ABCDEF", [][]int{{11, 15}, {4, 6}, {1, 2}, {3, 4}}, []string{"aaa", "bbb", "ccc", "ddd"})
	FailOnErr("%v", err)
	fPln(s) //         0CCC2*****BBB6789ATTTF
	//                 0CCC2*****BBB6789ATTTF
}

func TestProjectV(t *testing.T) {
	ss := []string{
		"12@a@1~b",
		"b~c",
		"c~a",
		"a~b~c",
		"b~c~d~e",
		"c~a",
		"d~e~b~a~f",
		"a",
		"***@b@***",
	}
	fPln(ProjectV(ss, "~", "", ""))
}
