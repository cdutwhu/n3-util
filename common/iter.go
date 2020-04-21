package common

import eg "github.com/cdutwhu/json-util/n3errs"

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
		FailOnErr("%v: params' count is [1,2,3]", eg.PARAM_INVALID)
	}
	if end <= start {
		FailOnErr("%v: [end](%d) must be greater than [start](%d)", eg.PARAM_INVALID, end, start)
	}
	if step < 1 {
		FailOnErr("%v: [step](%d) must be greater than 0", eg.PARAM_INVALID, step)
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
