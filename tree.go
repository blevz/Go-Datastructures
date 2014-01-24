package ds

import (
	"fmt"
)

type Tree struct {
	left  *Tree
	right *Tree
	val   interface{}
}

//Tree making methods

func MakeTree() (t Tree) {
	return
}

//With val

func MakeTreeWithVal(valtoset interface{}) (t Tree) {
	t.val = valtoset
	return
}

//With Subtrees

func MakeTreeWithLeftSubTree(lefty *Tree) (t Tree) {
	t.left = lefty
	return
}

func MakeTreeWithRightSubTree(righty *Tree) (t Tree) {
	t.right = righty
	return
}
func MakeTreeWithSubtrees(lefty *Tree, righty *Tree) (t Tree) {
	t.right = righty
	t.left = lefty
	return
}

func (t Tree) IsLeaf() bool {
	return (t.left == nil && t.right == nil)
}

func (t Tree) Depth() (d int64) {
	if t.IsLeaf() {
		d = 1
	} else {
		if t.left == nil {
			d = t.right.Depth() + 1
		} else if t.right == nil {
			d = t.left.Depth() + 1
		} else {
			ld := t.left.Depth() + 1
			rd := t.right.Depth() + 1

			if ld > rd {
				d = ld
			} else {
				d = rd
			}
		}
	}

	return
}

// Deletes

func (t *Tree) DeleteLeftSubtree() {
	t.left = nil
}

func (t *Tree) DeleteRightSubtree() {
	t.right = nil
}

func (t *Tree) DeleteLeftChild() {

}

func (t *Tree) DeleteRightChild() {

}

// Do's

func (t Tree) prefixDo(f func(interface{})) {
	f(t.val)
	if t.left != nil {
		t.left.prefixDo(f)
	}
	if t.right != nil {
		t.right.prefixDo(f)
	}
}

func (t Tree) inOrderDo(f func(interface{})) {
	if t.left != nil {
		t.left.prefixDo(f)
	}
	f(t.val)
	if t.right != nil {
		t.right.prefixDo(f)
	}
}

func (t Tree) postfixDo(f func(interface{})) {
	if t.left != nil {
		t.left.prefixDo(f)
	}
	if t.right != nil {
		t.right.prefixDo(f)
	}
	f(t.val)
}

//Prints

func (t Tree) prefixPrint() {
	t.prefixDo(func(x interface{}) {
		fmt.Println(x)
	})
}
