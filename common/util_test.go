package common

import (
	"testing"
)

func TestFailLog(t *testing.T) {
	logfile := "./log.txt"
	SetLog(logfile)
	FailOnErr("aa ")

	logfile = "./log1.txt"
	SetLog(logfile)

	logfile = "./log1.txt"
	SetLog(logfile)
	// FailOnErr("%v", fEf("test panic"))
}

func TestIter(t *testing.T) {
	for i := range N(10) {
		fPln(i)
	}
	fPln(" ------------ ")
	for i := range Iter(2, 3, 10) {
		fPln(i)
	}
	fPln(" ------------ ")
	func(slc ...int) {
		for _, a := range slc {
			fPln(a)
		}
	}(Iter2Slc(2, 3, 30)...)
}

func TestSHA(t *testing.T) {
	fPln("MD5", MD5Str("a"))       // 0cc175b9c0f1b6a831c399e269772661
	fPln("SHA1", SHA1Str("a"))     // 86f7e437faa5a7fce15d1ddcb9eaeaea377667b8
	fPln("SHA256", SHA256Str("a")) // ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
}

func TestIsSetCover(t *testing.T) {
	arr1 := []string{"a", "B", "c", "d"}
	arr2 := []string{"a", "b", "c"}
	fPln(CanSetCover(arr1, arr2))
	arr1 = []string{"c", "b", "a"}
	arr2 = []string{"a", "b", "c"}
	fPln(CanSetCover(arr1, arr2))
	// arr3 := 6
	// arr4 := 7
	// fPln(CanSetCover(arr3, arr4))
}

func TestToSet(t *testing.T) {
	fPln(ToSet([]int{1, 3, 2, 1, 3, 5}))
	fPln(ToSet([]string{"1", "2", "3", "4", "1", "3", "2"}))
}

func TestSetIntersect(t *testing.T) {
	arr1 := []string{"A", "B", "c", "d"}
	arr2 := []string{"C", "d", "a"}
	fPln(SetIntersect(arr1, arr2).([]string))
}

func TestSetUnion(t *testing.T) {
	arr1 := []string{"A", "B", "c", "d"}
	arr2 := []string{"C", "d", "A"}
	fPln(SetUnion(arr1, arr2).([]string))
}

func TestEncrypt(t *testing.T) {
	data := Encrypt([]byte("abc111"), "A")
	// fPln(string(data))
	bytes, _ := Decrypt(data, "A")
	fPln(string(bytes))
}

func TestMapPrint(t *testing.T) {
	MapPrint(map[string]string{
		// "a": "b",
		"3": "4 a",
		"5": "b sss",
		"7": "B    sss",
		"1": "2  2345678  223",
	})
}

func TestRmTailFromLast(t *testing.T) {
	fPln(RmTailFromLast("AB.CD.EF", "."))
	fPln(RmTailFromLast("AB.CD.EF", "#"))
}

func TestRmTailFromLastN(t *testing.T) {
	fPln(RmTailFromLastN("AB.CD.EF", ".", 2))
	fPln(RmTailFromLastN("AB.CD.EF", "#", 2))
}

func TestRmHeadToLast(t *testing.T) {
	fPln(RmHeadToLast("##AB##CD##F", "#"))
}

func TestRmHeadToFirst(t *testing.T) {
	fPln(RmHeadToFirst("777##AB##CD##F", "##"))
}

func TestColor(t *testing.T) {
	fPln("\033[31mRed")
	fPln("\033[32mGreen")
	fPln("\033[34mBlue")
}
