package ds

import (
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
)

func TestDset(t *testing.T) {
	Convey("Test Find", t, func() {
		var a Disjoint_Set_I
		a = MakeDJSet(3)
		So(a.Find(0), ShouldNotEqual, a.Find(1))
		So(a.Find(0), ShouldNotEqual, a.Find(2))
		So(a.Find(1), ShouldNotEqual, a.Find(2))
	})

	Convey("Test Union", t, func() {
		var a Disjoint_Set_I
		a = MakeDJSet(4)
		So(a.Find(0), ShouldNotEqual, a.Find(1))
		So(a.Find(0), ShouldNotEqual, a.Find(2))
		So(a.Find(0), ShouldNotEqual, a.Find(3))
		So(a.Find(1), ShouldNotEqual, a.Find(2))
		So(a.Find(1), ShouldNotEqual, a.Find(3))
		So(a.Find(2), ShouldNotEqual, a.Find(3))

		a.Union(0, 0)
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

		a.Union(2, 2)
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

func BenchmarkFind(b *testing.B) {
	var a Disjoint_Set_I
	a = MakeDJSet(b.N)
	for i := 0; i < b.N; i++ {
		a.Find(rand.Intn(b.N))
	}
}

func BenchmarkUnion(b *testing.B) {
	var a Disjoint_Set_I
	a = MakeDJSet(b.N)
	for i := 0; i < b.N; i++ {
		a.Union(rand.Intn(b.N), rand.Intn(b.N))
	}
}
