package common

import (
	"reflect"
	"regexp"
	"sort"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// MapKeys : only apply to single type key
func MapKeys(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	FailOnErrWhen(v.Kind() != reflect.Map, "%v", eg.PARAM_INVALID_MAP)
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := reflect.TypeOf(keys[0].Interface())
		rstValue := reflect.MakeSlice(reflect.SliceOf(kType), L, L)
		for i, k := range keys {
			rstValue.Index(i).Set(reflect.ValueOf(k.Interface()))
		}
		// sort keys if keys are int or float64 or string
		rst := rstValue.Interface()
		switch keys[0].Interface().(type) {
		case int:
			sort.Ints(rst.([]int))
		case float64:
			sort.Float64s(rst.([]float64))
		case string:
			sort.Strings(rst.([]string))
		}
		return rst
	}
	return nil
}

// MapKVs : only apply to single type key and single type value
func MapKVs(m interface{}) (interface{}, interface{}) {
	v := reflect.ValueOf(m)
	FailOnErrWhen(v.Kind() != reflect.Map, "%v", eg.PARAM_INVALID_MAP)
	keys := v.MapKeys()
	if L := len(keys); L > 0 {
		kType := reflect.TypeOf(keys[0].Interface())
		kRst := reflect.MakeSlice(reflect.SliceOf(kType), L, L)
		vType := reflect.TypeOf(v.MapIndex(keys[0]).Interface())
		vRst := reflect.MakeSlice(reflect.SliceOf(vType), L, L)
		for i, k := range keys {
			kRst.Index(i).Set(reflect.ValueOf(k.Interface()))
			vRst.Index(i).Set(reflect.ValueOf(v.MapIndex(k).Interface()))
		}
		return kRst.Interface(), vRst.Interface()
	}
	return nil, nil
}

// MapsJoin : overwritted by the 2nd params
func MapsJoin(m1, m2 interface{}) interface{} {
	v1, v2 := reflect.ValueOf(m1), reflect.ValueOf(m2)
	FailOnErrWhen(v1.Kind() != reflect.Map, "%v: m1", eg.MAP_INVALID)
	FailOnErrWhen(v2.Kind() != reflect.Map, "%v: m2", eg.MAP_INVALID)
	keys1, keys2 := v1.MapKeys(), v2.MapKeys()
	if len(keys1) > 0 && len(keys2) > 0 {
		k1, k2 := keys1[0], keys2[0]
		k1Type, k2Type := reflect.TypeOf(k1.Interface()), reflect.TypeOf(k2.Interface())
		v1Type, v2Type := reflect.TypeOf(v1.MapIndex(k1).Interface()), reflect.TypeOf(v2.MapIndex(k2).Interface())
		FailOnErrWhen(k1Type != k2Type, "%v", eg.MAPS_DIF_KEY_TYPE)
		FailOnErrWhen(v1Type != v2Type, "%v", eg.MAPS_DIF_VALUE_TYPE)
		aMap := reflect.MakeMap(reflect.MapOf(k1Type, v1Type))
		for _, k := range keys1 {
			aMap.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(v1.MapIndex(k).Interface()))
		}
		for _, k := range keys2 {
			aMap.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(v2.MapIndex(k).Interface()))
		}
		return aMap.Interface()
	}
	if len(keys1) > 0 && len(keys2) == 0 {
		return m1
	}
	if len(keys1) == 0 && len(keys2) > 0 {
		return m2
	}
	return m1
}

// MapsMerge : overwritted by the later params
func MapsMerge(ms ...interface{}) interface{} {
	if len(ms) == 0 {
		return nil
	}
	mm := ms[0]
	for i, m := range ms {
		if i >= 1 {
			mm = MapsJoin(mm, m)
		}
	}
	return mm
}

// MapPrint : Key Sorted Print
func MapPrint(m interface{}) {
	re := regexp.MustCompile(`^[+-]?[0-9]*\.?[0-9]+:`)
	mapstr := fSp(m)
	mapstr = mapstr[4 : len(mapstr)-1]
	fPln(mapstr)
	I := 0
	rmIdxList := []int{}
	ss := sSplit(mapstr, " ")
	for i, s := range ss {
		if re.MatchString(s) {
			I = i
		} else {
			ss[I] += " " + s
			rmIdxList = append(rmIdxList, i) // to be deleted (i)
		}
	}
	for i, s := range ss {
		if !XIn(i, rmIdxList) {
			fPln(i, s)
		}
	}
}
