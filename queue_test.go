package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Queue(t *testing.T) {
	Convey("Test Queue", t, func() {
		a := MakeQueue()
		So(a.Size(), ShouldEqual, 0)

		a.Push(1)
		So(a.Size(), ShouldEqual, 1)
		So(a.Front(), ShouldEqual, 1)

		a.Push(2)
		So(a.Size(), ShouldEqual, 2)
		So(a.Front(), ShouldEqual, 1)

		a.Push(3)
		So(a.Size(), ShouldEqual, 3)
		So(a.Front(), ShouldEqual, 1)

		a.Pop()
		So(a.Size(), ShouldEqual, 2)
		So(a.Front(), ShouldEqual, 2)

		a.Pop()
		So(a.Size(), ShouldEqual, 1)
		So(a.Front(), ShouldEqual, 3)

		a.Pop()
		So(a.Size(), ShouldEqual, 0)

		a.Push(5)
		a.Push(4)
		a.Push(3)
		a.Push(2)
		a.Push(1)
		So(a.Front(), ShouldEqual, 5)
		a.Pop()
		a.Push(2)
		a.Push(1)
		So(a.Front(), ShouldEqual, 4)
		a.Pop()
		So(a.Front(), ShouldEqual, 3)
	})
}
