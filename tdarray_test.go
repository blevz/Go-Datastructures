package ds

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTwoDArray(t *testing.T) {
	x := makeTDArray(5, 5)
	y := makeTDArray(2, 5)
	z := makeTDArray(3, 3)
	Convey("test creation size and get size", t, func() {
		a, b := x.GetSize()
		So(len(x.data), ShouldEqual, 25)
		So(a, ShouldEqual, 5)
		So(b, ShouldEqual, 5)

		a, b = y.GetSize()
		So(len(y.data), ShouldEqual, 10)
		So(a, ShouldEqual, 2)
		So(b, ShouldEqual, 5)

		a, b = z.GetSize()
		So(len(z.data), ShouldEqual, 9)
		So(a, ShouldEqual, 3)
		So(b, ShouldEqual, 3)

	})
	Convey("test bounds", t, func() {
		So(x.checkBounds(5, 5), ShouldBeFalse)
		So(x.checkBounds(4, 5), ShouldBeFalse)
		So(x.checkBounds(4, 4), ShouldBeTrue)

		So(y.checkBounds(2, 5), ShouldBeFalse)
		So(y.checkBounds(1, 5), ShouldBeFalse)
		So(y.checkBounds(0, 5), ShouldBeFalse)
		So(y.checkBounds(2, 4), ShouldBeFalse)
		So(y.checkBounds(1, 4), ShouldBeTrue)
	})

	Convey("test get", t, func() {
		x.data[0] = 1
		x.data[1] = 2
		x.data[2] = 3
		x.data[3] = 4
		x.data[4] = 5
		x.data[5] = 6
		x.data[6] = 7

		So(x.getElemVal(0, 0), ShouldEqual, 1)
		So(x.getElemVal(0, 1), ShouldEqual, 2)
		So(x.getElemVal(0, 2), ShouldEqual, 3)
		So(x.getElemVal(0, 3), ShouldEqual, 4)
		So(x.getElemVal(0, 4), ShouldEqual, 5)
		So(x.getElemVal(1, 0), ShouldEqual, 6)
		So(x.getElemVal(1, 1), ShouldEqual, 7)
	})
	Convey("test set", t, func() {
		x.setElem(0, 0, 2)
		x.setElem(0, 1, -1)
		x.setElem(0, 2, 55)
		x.setElem(1, 0, 22)
		So(x.getElemVal(0, 0), ShouldEqual, 2)
		So(x.getElemVal(0, 1), ShouldEqual, -1)
		So(x.getElemVal(0, 2), ShouldEqual, 55)
		So(x.getElemVal(1, 0), ShouldEqual, 22)

	})
}
