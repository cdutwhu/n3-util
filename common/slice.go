package common

import (
	"math"
	"reflect"

	eg "github.com/cdutwhu/n3-util/n3errs"
)

// SliceAttach :
func SliceAttach(s1, s2 interface{}, pos int) (interface{}, error) {
	v1, v2 := reflect.ValueOf(s1), reflect.ValueOf(s2)
	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		return nil, eg.PARAM_INVALID_SLICE
	}

	l1, l2 := v1.Len(), v2.Len()
	if l1 > 0 && l2 > 0 {
		if pos > l1 {
			return s1, nil
		}
		lm := int(math.Max(float64(l1), float64(l2+pos)))
		v := reflect.AppendSlice(v1.Slice(0, pos), v2)
		return v.Slice(0, lm).Interface(), nil
	}
	if l1 > 0 && l2 == 0 {
		return s1, nil
	}
	if l1 == 0 && l2 > 0 {
		return s2, nil
	}
	return s1, nil
}

// SliceCover :
func SliceCover(ss ...interface{}) interface{} {
	if len(ss) == 0 {
		return nil
	}
	attached := ss[0]
	for i, s := range ss {
		if i >= 1 {
			var err error
			attached, err = SliceAttach(attached, s, 0)
			FailOnErr("%v: ", err)
		}
	}
	return attached
}

// CanSetCover : check if setA contains setB ? return the first B-Index of which item is not in setA
func CanSetCover(setA, setB interface{}) (bool, int) {
	tA, tB := reflect.TypeOf(setA), reflect.TypeOf(setB)
	if tA != tB || (tA.Kind() != reflect.Slice && tA.Kind() != reflect.Array) {
		FailOnErr("%v: only can be [slice, array]", eg.PARAM_INVALID)
	}
	vA, vB := reflect.ValueOf(setA), reflect.ValueOf(setB)
	if vA.Len() < vB.Len() {
		return false, -1
	}
NEXT:
	for j := 0; j < vB.Len(); j++ {
		b := vB.Index(j).Interface()
		for i := 0; i < vA.Len(); i++ {
			if reflect.DeepEqual(b, vA.Index(i).Interface()) {
				continue NEXT
			}
			if i == vA.Len()-1 { // if b falls down to the last vA item position, which means vA doesn't have b item, return false
				return false, j
			}
		}
	}
	return true, -1
}

// SetIntersect :
func SetIntersect(setA, setB interface{}) interface{} {
	if setA == nil || setB == nil {
		return nil
	}

	tA, tB := reflect.TypeOf(setA), reflect.TypeOf(setB)
	if tA != tB || (tA.Kind() != reflect.Slice && tA.Kind() != reflect.Array) {
		FailOnErr("%v: only can be [slice, array]", eg.PARAM_INVALID)
	}
	vA, vB := reflect.ValueOf(setA), reflect.ValueOf(setB)
	set := reflect.MakeSlice(tA, 0, vA.Len())
NEXT:
	for j := 0; j < vB.Len(); j++ {
		b := vB.Index(j)
		for i := 0; i < vA.Len(); i++ {
			if reflect.DeepEqual(b.Interface(), vA.Index(i).Interface()) {
				set = reflect.Append(set, b)
				continue NEXT
			}
		}
	}
	return set.Interface()
}

// SetUnion :
func SetUnion(setA, setB interface{}) interface{} {
	switch {
	case setA != nil && setB == nil:
		return setA
	case setA == nil && setB != nil:
		return setB
	case setA == nil && setB == nil:
		return nil
	}

	tA, tB := reflect.TypeOf(setA), reflect.TypeOf(setB)
	if tA != tB || (tA.Kind() != reflect.Slice && tA.Kind() != reflect.Array) {
		FailOnErr("%v: only can be [slice, array]", eg.PARAM_INVALID)
	}
	vA, vB := reflect.ValueOf(setA), reflect.ValueOf(setB)
	set := reflect.MakeSlice(tA, 0, vA.Len()+vB.Len())
	set = reflect.AppendSlice(set, vA)
	set = reflect.AppendSlice(set, vB)
	return ToSet(set.Interface())
}

// ToSet : convert slice / array to set. i.e. remove duplicated items
func ToSet(slc interface{}) interface{} {
	if slc == nil {
		return nil
	}

	t := reflect.TypeOf(slc)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		FailOnErr("%v: only can be [slice, array]", eg.PARAM_INVALID)
	}
	v := reflect.ValueOf(slc)
	if v.Len() == 0 {
		return slc
	}

	set := reflect.MakeSlice(t, 0, v.Len())
	set = reflect.Append(set, v.Index(0))
NEXT:
	for i := 1; i < v.Len(); i++ {
		vItem := v.Index(i)
		for j := 0; j < set.Len(); j++ {
			if reflect.DeepEqual(vItem.Interface(), set.Index(j).Interface()) {
				continue NEXT
			}
			if j == set.Len()-1 { // if vItem falls down to the last set position, which means set doesn't have this item, then add it.
				set = reflect.Append(set, vItem)
			}
		}
	}
	return set.Interface()
}

// ToGeneralSlc :
func ToGeneralSlc(slc interface{}) (gslc []interface{}) {
	if slc == nil {
		return nil
	}

	v := reflect.ValueOf(slc)
	k := v.Type().Kind()
	if k != reflect.Slice && k != reflect.Array {
		FailOnErr("%v: only can be [slice, array]", eg.PARAM_INVALID)
	}
	for i := 0; i < v.Len(); i++ {
		gslc = append(gslc, v.Index(i).Interface())
	}
	return
}
