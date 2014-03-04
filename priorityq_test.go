package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := PriorityQ{}
	pq.Init()
	Convey("Test PQ", t, func() {
		pq.PushWithPriority("a", -3)
		pq.PushWithPriority("d", 1)
		pq.PushWithPriority("c", 0.0)
		pq.PushWithPriority("b", -1)
		pq.PushWithPriority("e", 3)

		v := pq.Pop()
		So(v, ShouldEqual, "a")
		v = pq.Pop()
		So(v, ShouldEqual, "b")
		v = pq.Pop()
		So(v, ShouldEqual, "c")
		v = pq.Pop()
		So(v, ShouldEqual, "d")
		v = pq.Pop()
		So(v, ShouldEqual, "e")
	})

	Convey("Test PQ2", t, func() {

		pq.PushWithPriority("b", -1)
		pq.PushWithPriority("a", -3)
		pq.PushWithPriority("e", 3)
		pq.PushWithPriority("d", 1)
		pq.PushWithPriority("c", 0.0)

		v := pq.Pop()
		So(v, ShouldEqual, "a")
		v = pq.Pop()
		So(v, ShouldEqual, "b")
		v = pq.Pop()
		So(v, ShouldEqual, "c")
		v = pq.Pop()
		So(v, ShouldEqual, "d")
		v = pq.Pop()
		So(v, ShouldEqual, "e")
	})
}
