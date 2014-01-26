package ds

//Some math helper functions for those pesky ints

func MaxInt(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func PowInt(a, b int64) (r int64) {

	if b == 0 {
		r = 1
	} else if b > 0 {
		r = a
		for x := int64(0); x < b; x++ {
			r *= a
		}
	} else if b < 0 {
		r = -1
	}
	return
}
