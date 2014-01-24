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

func TestTrees(t *testing.T) {
	Convey("Test Trees", t, func() {
		a := MakeTree()
		b := MakeTreeWithLeftSubTree(&a)
		e := MakeTree()
		d := MakeTreeWithRightSubTree(&e)
		c := MakeTreeWithSubtrees(&d, &b)

		a.val = 5
		b.val = "hello"
		c.val = "美国"
		d.val = 2.5
		e.val = false

		c.prefixPrint()
		So(a.IsLeaf(), ShouldBeTrue)
		So(b.IsLeaf(), ShouldBeFalse)

	})

	Convey("Test Prefix do", t, func() {
		a := MakeTree()
		b := MakeTreeWithLeftSubTree(&a)
		e := MakeTree()
		d := MakeTreeWithRightSubTree(&e)
		c := MakeTreeWithSubtrees(&d, &b)

		a.val = 5
		b.val = "hello"
		c.val = "美国"
		d.val = 2.5
		e.val = false

		c.prefixDo(func(x interface{}) {
			fmt.Println(x)
		})
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

func TestTree2(t *testing.T) {
	l1 := MakeTree()
	l2 := MakeTree()
	l3 := MakeTree()
	l4 := MakeTree()
	l5 := MakeTreeWithRightSubTree(&l2)

	t6 := MakeTreeWithLeftSubTree(&l3)
	t1 := MakeTreeWithSubtrees(&l1, &l5)
	t2 := MakeTreeWithSubtrees(&t6, &l4)
	t3 := MakeTreeWithLeftSubTree(&t1)
	t4 := MakeTreeWithRightSubTree(&t2)
	t5 := MakeTreeWithSubtrees(&t3, &t4)

	Convey("Test Trees", t, func() {

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

	Convey("test depth", t, func() {
		So(t5.Depth(), ShouldEqual, 5)
		So(t4.Depth(), ShouldEqual, 4)
		So(t3.Depth(), ShouldEqual, 4)
		So(t2.Depth(), ShouldEqual, 3)
		So(t1.Depth(), ShouldEqual, 3)

	})
}
