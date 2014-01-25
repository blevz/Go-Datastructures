package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDset(t *testing.T) {
	Convey("Test Find", t, func() {
		a := MakeDSet(3)
		So(a.Find(0), ShouldNotEqual, a.Find(1))
		So(a.Find(0), ShouldNotEqual, a.Find(2))
		So(a.Find(1), ShouldNotEqual, a.Find(2))
	})

	Convey("Test Union", t, func() {
		a := MakeDSet(4)
		So(a.Find(0), ShouldNotEqual, a.Find(1))
		So(a.Find(0), ShouldNotEqual, a.Find(2))
		So(a.Find(0), ShouldNotEqual, a.Find(3))
		So(a.Find(1), ShouldNotEqual, a.Find(2))
		So(a.Find(1), ShouldNotEqual, a.Find(3))
		So(a.Find(2), ShouldNotEqual, a.Find(3))

		a.Union(0, 1)
		So(a.Find(0), ShouldEqual, a.Find(1))
		So(a.Find(0), ShouldNotEqual, a.Find(2))
		So(a.Find(0), ShouldNotEqual, a.Find(3))
		So(a.Find(1), ShouldNotEqual, a.Find(2))
		So(a.Find(1), ShouldNotEqual, a.Find(3))
		So(a.Find(2), ShouldNotEqual, a.Find(3))

		a.Union(2, 3)
		So(a.Find(0), ShouldEqual, a.Find(1))
		So(a.Find(0), ShouldNotEqual, a.Find(2))
		So(a.Find(0), ShouldNotEqual, a.Find(3))
		So(a.Find(1), ShouldNotEqual, a.Find(2))
		So(a.Find(1), ShouldNotEqual, a.Find(3))
		So(a.Find(2), ShouldEqual, a.Find(3))

		a.Union(3, 1)
		So(a.Find(0), ShouldEqual, a.Find(1))
		So(a.Find(0), ShouldEqual, a.Find(2))
		So(a.Find(0), ShouldEqual, a.Find(3))
		So(a.Find(1), ShouldEqual, a.Find(2))
		So(a.Find(1), ShouldEqual, a.Find(3))
		So(a.Find(2), ShouldEqual, a.Find(3))
	})
}
