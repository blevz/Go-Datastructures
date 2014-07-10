package ds

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//      C
//	   / \
//	  d	  B
//	   \ /
//     e a

func TestBTreeFreeTypedness(t *testing.T) {
	a := MakeBTree()
	b := MakeBTreeWithLeftSubBTree(&a)
	e := MakeBTree()
	d := MakeBTreeWithRightSubBTree(&e)
	c := MakeBTreeWithSubtrees(&d, &b, "c")
	a.val = "a"
	b.val = "b"
	c.val = "c"
	d.val = "d"
	e.val = "e"

	Convey("Test BTrees", t, func() {
		//c.prefixPrint()
		So(a.IsLeaf(), ShouldBeTrue)
		So(b.IsLeaf(), ShouldBeFalse)

	})

}

//             t5
//            /  \
//          t3     t4
//         /        \
//       t1         t2
//     /   \      /    \
//   l1    l5    t6    l4
//           \   /
//            l2 l3

func TestBTreeHelperFunctions(t *testing.T) {
	l1 := MakeBTree()
	l2 := MakeBTree()
	l3 := MakeBTree()
	l4 := MakeBTree()
	l5 := MakeBTreeWithRightSubBTree(&l2)

	t6 := MakeBTreeWithLeftSubBTree(&l3)
	t1 := MakeBTreeWithSubtrees(&l1, &l5, "t1")
	t2 := MakeBTreeWithSubtrees(&t6, &l4, "t2")
	t3 := MakeBTreeWithLeftSubBTree(&t1)
	t4 := MakeBTreeWithRightSubBTree(&t2)
	t5 := MakeBTreeWithSubtrees(&t3, &t4, "t5")

	Convey("Test IsLeaf()", t, func() {

		So(l1.IsLeaf(), ShouldBeTrue)
		So(l2.IsLeaf(), ShouldBeTrue)
		So(l3.IsLeaf(), ShouldBeTrue)
		So(l4.IsLeaf(), ShouldBeTrue)

		So(t1.IsLeaf(), ShouldBeFalse)
		So(t2.IsLeaf(), ShouldBeFalse)
		So(t3.IsLeaf(), ShouldBeFalse)
		So(t4.IsLeaf(), ShouldBeFalse)
		So(t5.IsLeaf(), ShouldBeFalse)
	})

	Convey("test Depth()", t, func() {
		So(t5.Depth(), ShouldEqual, 5)
		So(t4.Depth(), ShouldEqual, 4)
		So(t3.Depth(), ShouldEqual, 4)
		So(t2.Depth(), ShouldEqual, 3)
		So(t1.Depth(), ShouldEqual, 3)
	})

}
