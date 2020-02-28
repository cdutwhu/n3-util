package common

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
