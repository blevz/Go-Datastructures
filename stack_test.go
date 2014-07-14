package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Stack(t *testing.T) {
	Convey("Test Stacks", t, func() {
		var a Stack_I
		a = MakeStack()
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
		a.Push(2)
		So(a.Size(), ShouldEqual, 2)
		So(a.Top(), ShouldEqual, 2)
		a.Push(3)
		So(a.Size(), ShouldEqual, 3)
		So(a.Top(), ShouldEqual, 3)

		a.Pop()
		So(a.Size(), ShouldEqual, 2)

		var b Stack_I
		b = MakeStack()
		b.Pop()
		b.Pop()
		b.Pop()
		So(b.Size(), ShouldEqual, 0)
		So(b.Top(), ShouldEqual, 0)

	})
}

func BenchmarkStack(b *testing.B) {
	a := MakeStack()
	for i := 0; i < b.N; i++ {
		a.Push(i)
		a.Pop()
	}
}
