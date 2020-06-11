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
func XIn(e, s interface{}) (bool, error) {
	v := reflect.ValueOf(s)
	if v.Kind() != reflect.Slice {
		return false, eg.PARAM_INVALID_SLICE
	}

	l := v.Len()
	for i := 0; i < l; i++ {
		if v.Index(i).Interface() == e {
			return true, nil
		}
	}
	return false, nil
}

// MatchAssign : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
func MatchAssign(chkCasesValues ...interface{}) (interface{}, error) {
	l := len(chkCasesValues)
	if l < 4 || l%2 == 1 {
		return nil, eg.PARAM_INVALID
	}
	_, l1, l2 := 1, (l-1)/2, (l-1)/2
	check := chkCasesValues[0]
	cases := chkCasesValues[1 : 1+l1]
	values := chkCasesValues[1+l1 : 1+l1+l2]
	for i, c := range cases {
		if check == c {
			return values[i], nil
		}
	}
	return chkCasesValues[l-1], nil
}

// TryInvoke : func Name must be Exportable
func TryInvoke(st interface{}, name string, args ...interface{}) (rets []interface{}, ok bool, err error) {
	defer func() {
		if r := recover(); r != nil {
			rets, ok, err = nil, false, fEf("%v", r)
		}
	}()

	stVal := reflect.ValueOf(st)
	if stVal.Kind() != reflect.Ptr || stVal.Elem().Kind() != reflect.Struct {
		return nil, false, eg.PARAM_INVALID_STRUCT_PTR
	}

	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	if _, ok := stVal.Type().MethodByName(name); ok {
		for _, ret := range stVal.MethodByName(name).Call(inputs) {
			rets = append(rets, ret.Interface())
		}
		return rets, true, nil
	}
	return rets, false, nil
}

// MustInvokeWithMW :
func MustInvokeWithMW(st interface{}, name string, args ...interface{}) []interface{} {
	rets, ok, err := TryInvokeWithMW(st, name, args...)
	FailOnErr("%v: [%s]", err, name)
	FailOnErrWhen(!ok, "%v: No [%s]", eg.INTERNAL, name)
	return rets
}

// TryInvokeWithMW : func Name must be Exportable
func TryInvokeWithMW(st interface{}, name string, args ...interface{}) (rets []interface{}, ok bool, err error) {
	m, e := Struct2Map(st)
	FailOnErr("%v", e)

	for k, v := range m {
		// fPln(k, v)
		if k == "MW" || k == "MiddleWare" || k == "MIDDLEWARE" {
			if mMW, ok := v.(map[string]map[string][]interface{}); ok {
			NEXTFN:
				for fn, mCallerParams := range mMW {
					for _, caller := range []string{name, "$@"} {
						if params, ok := mCallerParams[caller]; ok {
							// "$1" -> args[0] etc...
							for i, param := range params {
								if paramStr, ok := param.(string); ok && repParam.MatchString(paramStr) {
									num, err := scParseUint(paramStr[1:], 10, 64)
									FailOnErr("%v", err)
									FailOnErrWhen(int(num) > len(args), "MiddleWare: %v", eg.PARAM_INVALID_INDEX)
									params[i] = args[num-1]
								}
							}
							_, _, err = TryInvoke(st, fn, params...)
							FailOnErr("MiddleWare: %v", err)
							continue NEXTFN
						}
					}
				}
			}
		}
	}
	return TryInvoke(st, name, args...)
}

// InvRst :
func InvRst(rets interface{}, idx int) (ret interface{}, err error) {
	slc, ok := rets.([]interface{})
	if !ok {
		return nil, eg.PARAM_INVALID
	}
	if idx >= len(slc) {
		return nil, eg.PARAM_INVALID_INDEX
	}
	return slc[idx], nil
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
