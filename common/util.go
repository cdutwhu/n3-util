package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"time"
)

// IsNumeric :
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// N : for i := range N()
func N(n int) []struct{} {
	return make([]struct{}, n)
}

// Iter : for i := range Iter()
func Iter(params ...int) <-chan int {
	start, end, step := 0, 0, 1
	switch len(params) {
	case 1:
		end = params[0]
	case 2:
		start, end = params[0], params[1]
	case 3:
		start, step, end = params[0], params[1], params[2]
	default:
		FailOnErr("invalid params %v", fEf("@Iter"))
	}
	if end <= start {
		FailOnErr("[end](%d) must be greater than [start](%d) %v", end, start, fEf("@Iter"))
	}
	if step < 1 {
		FailOnErr("[step](%d) must be greater than 0 %v", step, fEf("@Iter"))
	}

	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i < end; i += step {
			ch <- i
		}
	}()
	return ch
}

// Iter2Slc :
func Iter2Slc(params ...int) (slc []int) {
	if len(params) == 1 {
		for i := range N(params[0]) {
			slc = append(slc, i)
		}
		return
	}
	for i := range Iter(params...) {
		slc = append(slc, i)
	}
	return
}

// SHA1Str :
func SHA1Str(s string) string {
	return fSf("%x", sha1.Sum([]byte(s)))
}

// SHA256Str :
func SHA256Str(s string) string {
	return fSf("%x", sha256.Sum256([]byte(s)))
}

// MD5Str :
func MD5Str(s string) string {
	return fSf("%x", md5.Sum([]byte(s)))
}

// TrackTime :
func TrackTime(start time.Time) {
	elapsed := time.Since(start)
	fPf("Took %s\n", elapsed)
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

// FuncTrack : full path of func name
// func FuncTrack(i interface{}) string {
// 	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
// }

// trackCaller :
func trackCaller() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(3, pc) // 3 is for util-FailLog. 2 is for "trackCaller" caller
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fSf("\n%s:%d\n%s\n", frame.File, frame.Line, frame.Function)
}

var log2file = false
var mPathFile map[string]*os.File = make(map[string]*os.File)

// SetLog :
func SetLog(logpath string) {
	if abspath, err := filepath.Abs(logpath); err == nil {
		if f, ok := mPathFile[abspath]; ok {
			log.SetFlags(log.LstdFlags)
			log.SetOutput(f)
			log2file = true
			return
		}
		if _, err := os.Stat(abspath); err == nil || os.IsNotExist(err) {
			if f, err := os.OpenFile(abspath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
				mPathFile[abspath] = f
				log.SetFlags(log.LstdFlags)
				log.SetOutput(f)
				log2file = true
			}
		}
	}
}

// ResetLog : call once at the exit
func ResetLog() {
	for logPath, f := range mPathFile {
		// delete empty error log
		fi, err := f.Stat()
		FailOnErr("%v", err)
		if fi.Size() == 0 {
			FailOnErr("%v", os.Remove(logPath))
		}
		// close
		f.Close()
	}
	mPathFile = make(map[string]*os.File)
	log.SetOutput(os.Stdout)
	log2file = false
}

// FailOnErr : error holder use "%v"
func FailOnErr(format string, v ...interface{}) {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					fatalInfo := fSf("FAIL: "+format+"%s\n", append(v, trackCaller())...)
					if log2file {
						fPln(time.Now().Format("2006/01/02 15:04:05 ") + fatalInfo)
					}
					log.Fatalf(fatalInfo)
				}
			}
		}
	}
}

// FailOnErrWhen :
func FailOnErrWhen(condition bool, format string, v ...interface{}) {
	if condition {
		for _, p := range v {
			switch p.(type) {
			case error:
				{
					if p != nil {
						fatalInfo := fSf("FAIL: "+format+"%s\n", append(v, trackCaller())...)
						if log2file {
							fPln(time.Now().Format("2006/01/02 15:04:05 ") + fatalInfo)
						}
						log.Fatalf(fatalInfo)
					}
				}
			}
		}
	}
}

// Log :
func Log(format string, v ...interface{}) (logItem string) {
	logItem = fSf("INFO: "+format+"%s\n", append(v, trackCaller())...)
	log.Printf("%s", logItem)
	if log2file {
		logItem = time.Now().Format("2006/01/02 15:04:05 ") + logItem
	}
	return
}

// LogWhen :
func LogWhen(condition bool, format string, v ...interface{}) (logItem string) {
	if condition {
		logItem = fSf("INFO: "+format+"%s\n", append(v, trackCaller())...)
		log.Printf("%s", logItem)
		if log2file {
			logItem = time.Now().Format("2006/01/02 15:04:05 ") + logItem
		}
	}
	return
}

// WarnOnErr :
func WarnOnErr(format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					log.Printf("WARN: "+format+"%s\n", append(v, trackCaller())...)
					return p.(error)
				}
			}
		}
	}
	return nil
}

// WarnOnErrWhen :
func WarnOnErrWhen(condition bool, format string, v ...interface{}) error {
	if condition {
		for _, p := range v {
			switch p.(type) {
			case error:
				{
					if p != nil {
						log.Printf("WARN: "+format+"%s\n", append(v, trackCaller())...)
						return p.(error)
					}
				}
			}
		}
	}
	return nil
}

// WrapOnErr :
func WrapOnErr(format string, v ...interface{}) error {
	for _, p := range v {
		switch p.(type) {
		case error:
			{
				if p != nil {
					return fmt.Errorf(format, v...)
				}
			}
		}
	}
	return nil
}

// -------------------------------------------------------- //

// Encrypt :
func Encrypt(data []byte, password string) []byte {
	if password == "" {
		return data
	}
	block, _ := aes.NewCipher([]byte(MD5Str(password)))
	gcm, err := cipher.NewGCM(block)
	FailOnErr("%v", err)
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	FailOnErr("%v", err)
	return gcm.Seal(nonce, nonce, data, nil)
}

// Decrypt :
func Decrypt(data []byte, password string) ([]byte, error) {
	if password == "" {
		return data, nil
	}
	key := []byte(MD5Str(password))
	block, err := aes.NewCipher(key)
	FailOnErr("%v", err)
	gcm, err := cipher.NewGCM(block)
	FailOnErr("%v", err)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	// FailOnErr("%v", err)
	return plaintext, err
}

// LocalIP returns the non loopback local IP of the host
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ""
}

// CanSetCover : check if setA contains setB ? return the first B-Index of which item is not in setA
func CanSetCover(setA, setB interface{}) (bool, int) {
	tA, tB := reflect.TypeOf(setA), reflect.TypeOf(setB)
	if tA != tB || (tA.Kind() != reflect.Slice && tA.Kind() != reflect.Array) {
		FailOnErr("%v", fEf("parameters only can be [slice] or [array]"))
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
	tA, tB := reflect.TypeOf(setA), reflect.TypeOf(setB)
	if tA != tB || (tA.Kind() != reflect.Slice && tA.Kind() != reflect.Array) {
		FailOnErr("%v", fEf("parameters only can be [slice] or [array]"))
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
	tA, tB := reflect.TypeOf(setA), reflect.TypeOf(setB)
	if tA != tB || (tA.Kind() != reflect.Slice && tA.Kind() != reflect.Array) {
		FailOnErr("%v", fEf("parameters only can be [slice] or [array]"))
	}
	vA, vB := reflect.ValueOf(setA), reflect.ValueOf(setB)
	set := reflect.MakeSlice(tA, 0, vA.Len()+vB.Len())
	set = reflect.AppendSlice(set, vA)
	set = reflect.AppendSlice(set, vB)
	return ToSet(set.Interface())
}

// ToSet : convert slice / array to set. i.e. remove duplicated items
func ToSet(slc interface{}) interface{} {
	t := reflect.TypeOf(slc)
	if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
		FailOnErr("%v", fEf("parameter only can be [slice] or [array]"))
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

// move from "cdutwhu/go-wrappers"

// HasAnyPrefix :
func HasAnyPrefix(s string, lsPrefix ...string) bool {
	FailOnErrWhen(len(lsPrefix) == 0, "%v", fEf("at least one prefix for input"))
	for _, prefix := range lsPrefix {
		if sHasPrefix(s, prefix) {
			return true
		}
	}
	return false
}

// RmTailFromLast :
func RmTailFromLast(s, mark string) string {
	if i := sLastIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

// RmTailFromLastN :
func RmTailFromLastN(s, mark string, iLast int) string {
	rt := s
	for i := 0; i < iLast; i++ {
		rt = RmTailFromLast(rt, mark)
	}
	return rt
}

// RmTailFromFirst :
func RmTailFromFirst(s, mark string) string {
	if i := sIndex(s, mark); i >= 0 {
		return s[:i]
	}
	return s
}

// RmTailFromFirstAny :
func RmTailFromFirstAny(s string, marks ...string) string {
	if len(marks) == 0 {
		return s
	}
	const MAX = 1000000
	var I int = MAX
	for _, mark := range marks {
		if i := sIndex(s, mark); i >= 0 && i < I {
			I = i
		}
	}
	if I != MAX {
		return s[:I]
	}
	return s
}

// RmHeadToLast :
func RmHeadToLast(s, mark string) string {
	if i := sLastIndex(s, mark); i >= 0 {
		return s[i+len(mark) : len(s)]
	}
	return s
}

// RmHeadToFirst :
func RmHeadToFirst(s, mark string) string {
	segs := sSpl(s, mark)
	if len(segs) > 1 {
		return sJoin(segs[1:], mark)
	}
	return s
}

// move from "cdutwhu/go-util"

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
	FailOnErrWhen(v.Kind() != reflect.Slice, "%v", fEf("s is NOT a SLICE!"))
	l := v.Len()
	for i := 0; i < l; i++ {
		if v.Index(i).Interface() == e {
			return true
		}
	}
	return false
}

// MapKeys :
func MapKeys(m interface{}) interface{} {
	v := reflect.ValueOf(m)
	FailOnErrWhen(v.Kind() != reflect.Map, "%v", fEf("NOT A MAP!"))
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

// MapKVs :
func MapKVs(m interface{}) (interface{}, interface{}) {
	v := reflect.ValueOf(m)
	FailOnErrWhen(v.Kind() != reflect.Map, "%v", fEf("NOT A MAP!"))
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
	FailOnErrWhen(v1.Kind() != reflect.Map, "%v", fEf("m1 is NOT A MAP!"))
	FailOnErrWhen(v2.Kind() != reflect.Map, "%v", fEf("m2 is NOT A MAP!"))
	keys1, keys2 := v1.MapKeys(), v2.MapKeys()
	if len(keys1) > 0 && len(keys2) > 0 {
		k1, k2 := keys1[0], keys2[0]
		k1Type, k2Type := reflect.TypeOf(k1.Interface()), reflect.TypeOf(k2.Interface())
		v1Type, v2Type := reflect.TypeOf(v1.MapIndex(k1).Interface()), reflect.TypeOf(v2.MapIndex(k2).Interface())
		FailOnErrWhen(k1Type != k2Type, "%v", fEf("different maps' key type!"))
		FailOnErrWhen(v1Type != v2Type, "%v", fEf("different maps' value type!"))
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
	ss := sSpl(mapstr, " ")
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

// SliceAttach :
func SliceAttach(s1, s2 interface{}, pos int) interface{} {
	v1, v2 := reflect.ValueOf(s1), reflect.ValueOf(s2)
	FailOnErrWhen(v1.Kind() != reflect.Slice, "%v", fEf("s1 is NOT a SLICE!"))
	FailOnErrWhen(v2.Kind() != reflect.Slice, "%v", fEf("s2 is NOT a SLICE!"))
	l1, l2 := v1.Len(), v2.Len()
	if l1 > 0 && l2 > 0 {
		if pos > l1 {
			return s1
		}
		lm := int(math.Max(float64(l1), float64(l2+pos)))
		v := reflect.AppendSlice(v1.Slice(0, pos), v2)
		return v.Slice(0, lm).Interface()
	}
	if l1 > 0 && l2 == 0 {
		return s1
	}
	if l1 == 0 && l2 > 0 {
		return s2
	}
	return s1
}

// SliceCover :
func SliceCover(ss ...interface{}) interface{} {
	if len(ss) == 0 {
		return nil
	}
	attached := ss[0]
	for i, s := range ss {
		if i >= 1 {
			attached = SliceAttach(attached, s, 0)
		}
	}
	return attached
}

// MatchAssign : NO ShortCut, MUST all valid, e.g. type assert, nil pointer, out of index
func MatchAssign(chkCasesValues ...interface{}) interface{} {
	l := len(chkCasesValues)
	FailOnErrWhen(l < 4 || l%2 == 1, "%v", fEf("Invalid parameters"))
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

// -------------------------------------------------------- //

// IsXML :
func IsXML(str string) bool {
	return xml.Unmarshal([]byte(str), new(interface{})) == nil
}

// XMLRoot :
func XMLRoot(xml string) (root string) {
	xml = sTrim(xml, " \t\n\r")
	FailOnErrWhen(!IsXML(xml), "%v", fEf("Invalid XML"))

	start, end := 0, 0
	for i := len(xml) - 1; i >= 0; i-- {
		switch xml[i] {
		case '>':
			end = i
		case '/':
			start = i + 1
		}
		if start != 0 && end != 0 {
			break
		}
	}
	root = xml[start:end]

	// check, flag (?s) let . includes "NewLine"
	// re1 := regexp.MustCompile(fSf(`(?s)^<%s .+</%s>$`, root, root))
	// re2 := regexp.MustCompile(fSf(`(?s)^<%s>.+</%s>$`, root, root))
	// FailOnErrWhen(!re1.MatchString(xml) && !re2.MatchString(xml), "%v", fEf("Invalid XML"))
	return
}

// IsJSON :
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// JSONRoot :
func JSONRoot(jsonstr string) string {
	x := make(map[string]interface{})
	FailOnErr("%v", json.Unmarshal([]byte(jsonstr), &x))
	for k := range x {
		return k
	}
	return ""
}

// -------------------------------------------------------- //

// ReplByPosGrp :
func ReplByPosGrp(s string, posGrp [][]int, newStrGrp []string) (ret string) {
	if len(posGrp) == 0 {
		return s
	}

	FailOnErrWhen(len(posGrp) != len(newStrGrp) && len(newStrGrp) != 1,
		"%v",
		fEf("posGrp's length must be equal to newStrGrp's length OR newStrGrp only has 1 element for filling all posGrp"))

	wrapper := make([]struct {
		posPair []int
		newStr  string
	}, len(posGrp))
	for i, pair := range posGrp {
		wrapper[i].posPair = pair
		if len(newStrGrp) == 1 {
			wrapper[i].newStr = newStrGrp[0]
		} else {
			wrapper[i].newStr = newStrGrp[i]
		}
	}
	sort.Slice(wrapper, func(i, j int) bool {
		return wrapper[i].posPair[0] < wrapper[j].posPair[0]
	})

	complement := make([][2]int, len(posGrp)+1)
	for i := 0; i < len(complement); i++ {
		if i == 0 {
			complement[i][0] = 0
			complement[i][1] = wrapper[i].posPair[0]
		} else if i == len(complement)-1 {
			complement[i][0] = wrapper[i-1].posPair[1]
			complement[i][1] = len(s)
		} else {
			complement[i][0] = wrapper[i-1].posPair[1]
			complement[i][1] = wrapper[i].posPair[0]
		}
	}

	keepStrGrp := make([]string, len(complement))
	for i := 0; i < len(keepStrGrp); i++ {
		start, end := complement[i][0], complement[i][1]
		FailOnErrWhen(end < start, "%v", fEf("end must be greater than start"))
		keepStrGrp[i] = s[start:end]
	}

	for i, keep := range keepStrGrp[:len(keepStrGrp)-1] {
		ret += (keep + wrapper[i].newStr)
	}
	ret += keepStrGrp[len(keepStrGrp)-1]

	return
}