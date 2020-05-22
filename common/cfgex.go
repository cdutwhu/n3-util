package common

import "reflect"

// CfgRepl :
func CfgRepl(cfg interface{}, mRepl map[string]interface{}) (interface{}, error) {
	if mRepl == nil || len(mRepl) == 0 {
		return cfg, nil
	}
	if reflect.ValueOf(cfg).Kind() == reflect.Ptr {
		if cfgElem := reflect.ValueOf(cfg).Elem(); cfgElem.Kind() == reflect.Struct {
			for i, nField := 0, cfgElem.NumField(); i < nField; i++ {
				for key, value := range mRepl {
					field := cfgElem.Field(i)
					if oriVal, ok := field.Interface().(string); ok && sContains(oriVal, key) {
						if repVal, ok := value.(string); ok {
							field.SetString(sReplaceAll(oriVal, key, repVal))
						} else {
							field.Set(reflect.ValueOf(value))
						}
					}
					// go into struct element
					if field.Kind() == reflect.Struct {
						for j, nFieldSub := 0, field.NumField(); j < nFieldSub; j++ {
							fieldSub := field.Field(j)
							if oriVal, ok := fieldSub.Interface().(string); ok && sContains(oriVal, key) {
								if repVal, ok := value.(string); ok {
									fieldSub.SetString(sReplaceAll(oriVal, key, repVal))
								} else {
									fieldSub.Set(reflect.ValueOf(value))
								}
							}
						}
					}
				}
			}
			return cfg, nil
		}
	}
	return nil, fEf("input cfg MUST be struct pointer")
}
