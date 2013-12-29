package ds

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSet(t *testing.T) {
	es := EmptySet()

	Convey("Empty set should have size zero and no nil pointers", t, func() {
		So(es, ShouldNotBeNil)
		So(es.size(), ShouldEqual, 0)
	})

	Convey("Test adding elements", t, func() {
		es.addElem("one")
		So(es.size(), ShouldEqual, 1)
		es.addElem("two")
		So(es.size(), ShouldEqual, 2)
		es.addElem("three")
		So(es.size(), ShouldEqual, 3)
	})

	Convey("Don't grow on duplicates", t, func() {
		es2 := EmptySet()
		es2.addElem("one")
		So(es2.size(), ShouldEqual, 1)
		es2.addElem("two")
		So(es2.size(), ShouldEqual, 2)
		es2.addElem("one")
		So(es2.size(), ShouldEqual, 2)
		es2.addElem("three")
		So(es2.size(), ShouldEqual, 3)
		es2.addElem("three")
		So(es2.size(), ShouldEqual, 3)
		es2.addElem("one")
		So(es2.size(), ShouldEqual, 3)
	})

	Convey("Test removing elements", t, func() {
		es3 := EmptySet()
		es3.addElem("one")
		es3.addElem("two")
		So(es3.size(), ShouldEqual, 2)
		es3.delElem("three")
		So(es3.size(), ShouldEqual, 2)
		es3.delElem("one")
		So(es3.size(), ShouldEqual, 1)
		es3.delElem("one")
		So(es3.size(), ShouldEqual, 1)
		es3.delElem("zero")
		So(es3.size(), ShouldEqual, 1)
	})

	es4 := EmptySet()
	es5 := EmptySet()
	es4.addElem("one")
	es4.addElem("two")
	es5.addElem("two")
	es5.addElem("five")
	es5.addElem("six")

	Convey("test exists", t, func() {
		So(es4.Exist("one"), ShouldBeTrue)
		So(es4.Exist("three"), ShouldBeFalse)
	})

	Convey("test intersect", t, func() {
		es6 := es4.intersect(es5)
		So(es6.size(), ShouldEqual, 1)
		So(es6.Exist("two"), ShouldBeTrue)
	})
	Convey("test union", t, func() {
		es6 := es5.union(es4)
		So(es6.size(), ShouldEqual, 4)
	})
	Convey("test difference", t, func() {
		es7 := es5.difference(es4)
		fmt.Println(es7.data)
		So(es7.size(), ShouldEqual, 2)
	})

}
