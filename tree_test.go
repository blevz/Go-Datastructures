package ds

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//      C
//	   / \
//	  d	  B
//	   \ /
//     e a

func TestBTreeFreeTypedness(t *testing.T) {
	a := MakeBinaryTree("a")
	b := MakeBinaryTreeWithLeftSubtree("b", a)
	e := MakeBinaryTree("e")
	d := MakeBinaryTreeWithRightSubtree("d", e)
	c := MakeBinaryTreeWithSubtrees("c", d, b)

	Convey("Test BTrees", t, func() {
		//c.prefixPrint()
		So(a.IsLeaf(), ShouldBeTrue)
		So(b.IsLeaf(), ShouldBeFalse)
		So(c.IsLeaf(), ShouldBeFalse)
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
	l1 := MakeBinaryTree(1)
	l2 := MakeBinaryTree(2)
	l3 := MakeBinaryTree(3)
	l4 := MakeBinaryTree(4)
	l5 := MakeBinaryTreeWithRightSubtree(5, l2)

	t6 := MakeBinaryTreeWithLeftSubtree(6, l3)
	t1 := MakeBinaryTreeWithSubtrees("t1", l1, l5)
	t2 := MakeBinaryTreeWithSubtrees("t2", t6, l4)
	t3 := MakeBinaryTreeWithLeftSubtree("t3", t1)
	t4 := MakeBinaryTreeWithRightSubtree("t4", t2)
	t5 := MakeBinaryTreeWithSubtrees("t5", t3, t4)

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
