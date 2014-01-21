package ds

type TDArray struct {
	data          []float64
	height, width int
}

func makeTDArray(h, w int) (t TDArray) {
	t.data = make([]float64, h*w)
	t.height = h
	t.width = w
	return
}

func (t TDArray) GetSize() (int, int) {
	return t.height, t.width
}

func (t TDArray) checkBounds(y, x int) bool {
	//Returns true if in bounds
	if (x >= 0 && x < t.width) && (y >= 0 && y < t.height) {
		return true
	} else {
		return false
	}
}

func (t TDArray) getElemVal(y, x int) float64 {
	if t.checkBounds(y, x) {
		return t.data[y*t.width+x]
	} else {
		panic("out of bounds")
	}
}

func (t *TDArray) setElem(y, x int, val float64) {
	if t.checkBounds(y, x) {
		t.data[y*t.width+x] = val
	}
}

func (t *TDArray) getElemRef(y, x int) *float64 {
	if t.checkBounds(y, x) {
		return &t.data[y*t.width+x]
	} else {
		return nil
	}
}
