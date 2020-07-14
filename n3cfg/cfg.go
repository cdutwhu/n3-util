package n3cfg

import (
	"net/url"
)

// WrapOnErr :
// func WrapOnErr(format string, v ...interface{}) error {
// 	for _, p := range v {
// 		switch p.(type) {
// 		case error:
// 			{
// 				if p != nil {
// 					return fmt.Errorf(format, v...)
// 				}
// 			}
// 		}
// 	}
// 	return nil
// }

// CfgRepl :
func CfgRepl(cfg interface{}, mRepl map[string]interface{}) interface{} {
	if mRepl == nil || len(mRepl) == 0 {
		return cfg
	}

	if vof(cfg).Kind() == typPTR {
		if cfgElem := vof(cfg).Elem(); cfgElem.Kind() == typSTRUCT {
			for i, nField := 0, cfgElem.NumField(); i < nField; i++ {
				for key, value := range mRepl {
					field := cfgElem.Field(i)
					if oriVal, ok := field.Interface().(string); ok && sContains(oriVal, key) {
						if repVal, ok := value.(string); ok {
							field.SetString(sReplaceAll(oriVal, key, repVal))
						} else {
							field.Set(vof(value))
						}
					}
					// go into struct element
					if field.Kind() == typSTRUCT {
						for j, nFieldSub := 0, field.NumField(); j < nFieldSub; j++ {
							fieldSub := field.Field(j)
							if oriVal, ok := fieldSub.Interface().(string); ok && sContains(oriVal, key) {
								if repVal, ok := value.(string); ok {
									fieldSub.SetString(sReplaceAll(oriVal, key, repVal))
								} else {
									fieldSub.Set(vof(value))
								}
							}
						}
					}
				}
			}
			return cfg
		}
	}
	failP1OnErr("%v", fEf("input cfg MUST be struct pointer"))
	return nil
}

// URLValues :
func URLValues(values url.Values, params ...string) (ok bool, lsValues [][]string) {
	for _, param := range params {
		if pv, ok := values[param]; ok {
			lsValues = append(lsValues, pv)
		}
	}
	if len(lsValues) == len(params) {
		return true, lsValues
	}
	return false, nil
}

// URLOneValueList :
// pick up one index-fixed value item from each values array, then combine them into one array
func URLOneValueList(values url.Values, idx int, params ...string) (ok bool, lsOneValue []string) {
	if ok, lsValues := URLValues(values, params...); ok {
		for _, vs := range lsValues {
			lsOneValue = append(lsOneValue, vs[idx])
		}
	}
	if len(params) == len(lsOneValue) {
		return true, lsOneValue
	}
	return false, nil
}

// URL1Value :
func URL1Value(values url.Values, idx int, params ...string) (bool, string) {
	if ok, ls1Value := URLOneValueList(values, idx, params...); ok {
		return true, ls1Value[0]
	}
	return false, ""
}

// URL2Values :
func URL2Values(values url.Values, idx int, params ...string) (bool, string, string) {
	if ok, ls2Values := URLOneValueList(values, idx, params...); ok {
		return true, ls2Values[0], ls2Values[1]
	}
	return false, "", ""
}

// URL3Values :
func URL3Values(values url.Values, idx int, params ...string) (bool, string, string, string) {
	if ok, ls3Values := URLOneValueList(values, idx, params...); ok {
		return true, ls3Values[0], ls3Values[1], ls3Values[2]
	}
	return false, "", "", ""
}

// URL4Values :
func URL4Values(values url.Values, idx int, params ...string) (bool, string, string, string, string) {
	if ok, ls4Values := URLOneValueList(values, idx, params...); ok {
		return true, ls4Values[0], ls4Values[1], ls4Values[2], ls4Values[3]
	}
	return false, "", "", "", ""
}

// URL5Values :
func URL5Values(values url.Values, idx int, params ...string) (bool, string, string, string, string, string) {
	if ok, ls5Values := URLOneValueList(values, idx, params...); ok {
		return true, ls5Values[0], ls5Values[1], ls5Values[2], ls5Values[3], ls5Values[4]
	}
	return false, "", "", "", "", ""
}

// URL6Values :
func URL6Values(values url.Values, idx int, params ...string) (bool, string, string, string, string, string, string) {
	if ok, ls6Values := URLOneValueList(values, idx, params...); ok {
		return true, ls6Values[0], ls6Values[1], ls6Values[2], ls6Values[3], ls6Values[4], ls6Values[5]
	}
	return false, "", "", "", "", "", ""
}
