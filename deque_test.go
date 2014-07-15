package ds

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_Reg_Deque(t *testing.T) {
	Convey("Test Deque", t, func() {
		var a Deque_I
		a = MakeDeque()
		So(a.Size(), ShouldEqual, 0)
		a.Push_Back(5)
		a.Push_Back(10)
		So(a.Front(), ShouldEqual, 5)
		So(a.Back(), ShouldEqual, 10)

		var b Deque_I
		b = MakeDeque()
		b.Push_Front(5)
		b.Push_Front(10)
		So(b.Front(), ShouldEqual, 10)
		So(b.Back(), ShouldEqual, 5)

		b.Push_Front(15)
		b.Push_Front(20)

		So(b.Front(), ShouldEqual, 20)
		So(b.Back(), ShouldEqual, 5)

		So(b.Pop_Front(), ShouldEqual, 20)
		So(b.Pop_Back(), ShouldEqual, 5)

		b.Print()

		So(b.Front(), ShouldEqual, 15)
		So(b.Back(), ShouldEqual, 10)

	})

}
