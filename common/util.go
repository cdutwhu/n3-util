package common

import (
	"reflect"
	"strconv"
	"time"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// IsNumeric :
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// TrackTime :
func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	fPf("Took %s\n", elapsed)
}

// IF : Ternary Operator LIKE < ? : >, BUT NO S/C, so block1 and block2 MUST all valid. e.g. type assert, nil pointer, out of index
func IF(condition bool, block1, block2 interface{}) interface{} {
	if condition {
		return block1
	}
	return block2
}

// XIn :
func XIn(e, s interface{}) bool {
	v := reflect.ValueOf(s)
	FailOnErrWhen(v.Kind() != reflect.Slice, "%v: [s]", eg.SLICE_INVALID)
	l := v.Len()
	for i := 0; i < l; i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

// MatchAssign : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
func MatchAssign(chkCasesValues ...interface{}) interface{} {
	l := len(chkCasesValues)
	FailOnErrWhen(l < 4 || l%2 == 1, "%v: ", eg.PARAM_INVALID)
	_, l1, l2 := 1, (l-1)/2, (l-1)/2
	check := chkCasesValues[0]
	cases := chkCasesValues[1 : 1+l1]
	values := chkCasesValues[1+l1 : 1+l1+l2]
	for i, c := range cases {
		if check == c {
			return values[i]
		}
	}
	return chkCasesValues[l-1]
}

// var (
// 	Color = func(colorString string) func(...interface{}) string {
// 		sprint := func(args ...interface{}) string {
// 			return fmt.Sprintf(colorString, fmt.Sprint(args...))
// 		}
// 		return sprint
// 	}
// 	Black   = Color("\033[1;30m%s\033[0m")
// 	Red     = Color("\033[1;31m%s\033[0m")
// 	Green   = Color("\033[1;32m%s\033[0m")
// 	Yellow  = Color("\033[1;33m%s\033[0m")
// 	Purple  = Color("\033[1;34m%s\033[0m")
// 	Magenta = Color("\033[1;35m%s\033[0m")
// 	Teal    = Color("\033[1;36m%s\033[0m")
// 	White   = Color("\033[1;37m%s\033[0m")
// )

// var (
// 	Info  = Teal
// 	Warn  = Yellow
// 	Fatal = Red
// )
