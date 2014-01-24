package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Stack(t *testing.T) {
	Convey("Test Stacks", t, func() {
		a := MakeStack()
		So(a.Size(), ShouldEqual, 0)
		a.Push(1)
		So(a.Size(), ShouldEqual, 1)
		So(a.Top(), ShouldEqual, 1)
		a.Push(2)
		So(a.Size(), ShouldEqual, 2)
		So(a.Top(), ShouldEqual, 2)
		a.Push(3)
		So(a.Size(), ShouldEqual, 3)
		So(a.Top(), ShouldEqual, 3)

		a.Pop()
		So(a.Size(), ShouldEqual, 2)
		So(a.Top(), ShouldEqual, 2)
		a.Pop()
		So(a.Size(), ShouldEqual, 1)
		So(a.Top(), ShouldEqual, 1)
		a.Pop()
		So(a.Size(), ShouldEqual, 0)
	})
}
