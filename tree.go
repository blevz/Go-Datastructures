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

//Make copy work

func CopyTreesToSubtrees(lefty Tree, righty Tree) (t Tree) {
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
	// TO DO
}

func (t *Tree) DeleteRightChild() {
	// TO DO
}

//Helper Functions

func (t Tree) IsLeaf() bool {
	return (t.left == nil && t.right == nil)
}

func (t Tree) HasLeftChild() bool {
	return t.left != nil
}

func (t Tree) HasRightChild() bool {
	return t.right != nil
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

func (t Tree) breadthFirstDo(f func(interface{})) {
	q := MakeQueue()
	q.Push(&t)

	for !q.Empty() {
		p := q.Front().(*Tree)
		q.Pop()
		f(p.val)

		//Enque Child nodes
		if p.HasLeftChild() {
			q.Push(p.left)
		}
		if p.HasRightChild() {
			q.Push(p.right)
		}
	}
}

func (t Tree) depthFirstDo(f func(interface{})) {
	s := MakeStack()
	s.Push(&t)

	for !s.Empty() {
		p := s.Top().(*Tree)
		s.Pop()
		f(p.val)

		//Enque Child nodes
		if p.HasLeftChild() {
			s.Push(p.left)
		}
		if p.HasRightChild() {
			s.Push(p.right)
		}
	}
}

//Prints

func (t Tree) prefixPrint() {
	t.prefixDo(func(x interface{}) {
		fmt.Println(x, " ")
	})
	fmt.Println("")
}

func (t Tree) postfixPrint() {
	t.postfixDo(func(x interface{}) {
		fmt.Println(x, " ")
	})
	fmt.Println("")
}

func (t Tree) inOrderPrint() {
	t.inOrderDo(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println("")
}

func (t Tree) BFPrint() {
	t.breadthFirstDo(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println("")
}

func (t Tree) DFPrint() {
	t.depthFirstDo(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println("")
}
