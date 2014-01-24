package ds

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSet(t *testing.T) {
	es := MakeEmptySet()

	Convey("Empty set should have Size zero and no nil pointers", t, func() {
		So(es, ShouldNotBeNil)
		So(es.Size(), ShouldEqual, 0)
	})

	Convey("Test adding elements", t, func() {
		es.AddElem("one")
		So(es.Size(), ShouldEqual, 1)
		es.AddElem("two")
		So(es.Size(), ShouldEqual, 2)
		es.AddElem("three")
		So(es.Size(), ShouldEqual, 3)
	})

	Convey("Don't grow on duplicates", t, func() {
		es2 := MakeEmptySet()
		es2.AddElem("one")
		So(es2.Size(), ShouldEqual, 1)
		es2.AddElem("two")
		So(es2.Size(), ShouldEqual, 2)
		es2.AddElem("one")
		So(es2.Size(), ShouldEqual, 2)
		es2.AddElem("three")
		So(es2.Size(), ShouldEqual, 3)
		es2.AddElem("three")
		So(es2.Size(), ShouldEqual, 3)
		es2.AddElem("one")
		So(es2.Size(), ShouldEqual, 3)
	})

	Convey("Test removing elements", t, func() {
		es3 := MakeEmptySet()
		es3.AddElem("one")
		es3.AddElem("two")
		So(es3.Size(), ShouldEqual, 2)
		es3.DelElem("three")
		So(es3.Size(), ShouldEqual, 2)
		es3.DelElem("one")
		So(es3.Size(), ShouldEqual, 1)
		es3.DelElem("one")
		So(es3.Size(), ShouldEqual, 1)
		es3.DelElem("zero")
		So(es3.Size(), ShouldEqual, 1)
	})

	es4 := MakeEmptySet()
	es5 := MakeEmptySet()
	es4.AddElem("one")
	es4.AddElem("two")
	es5.AddElem("two")
	es5.AddElem("five")
	es5.AddElem("six")

	Convey("test exists", t, func() {
		So(es4.Exist("one"), ShouldBeTrue)
		So(es4.Exist("three"), ShouldBeFalse)
	})

	Convey("test intersect", t, func() {
		es6 := es4.Intersect(es5)
		So(es6.Size(), ShouldEqual, 1)
		So(es6.Exist("two"), ShouldBeTrue)
	})
	Convey("test union", t, func() {
		es6 := es5.Union(es4)
		So(es6.Size(), ShouldEqual, 4)
	})
	Convey("test difference", t, func() {
		es7 := es5.Difference(es4)
		fmt.Println(es7.data)
		So(es7.Size(), ShouldEqual, 2)
	})

}
