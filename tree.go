package ds

import (
	"container/list"
	"fmt"
)

type BTree struct {
	left  *BTree
	right *BTree
	val   interface{}
}

type Tree struct {
	children list.List
	val      interface{}
}

//BTree making methods

func MakeBTree() (t BTree) {
	return
}

//With val

func MakeBTreeWithVal(valtoset interface{}) (t BTree) {
	t.val = valtoset
	return
}

//With Subtrees

func MakeBTreeWithLeftSubBTree(lefty *BTree) (t BTree) {
	t.left = lefty
	return
}

func MakeBTreeWithRightSubBTree(righty *BTree) (t BTree) {
	t.right = righty
	return
}
func MakeBTreeWithSubtrees(lefty *BTree, righty *BTree, val interface{}) (t BTree) {
	t.right = righty
	t.left = lefty
	t.val = val
	return
}

//Make copy work

func CopyBTreesToSubtrees(lefty BTree, righty BTree) (t BTree) {
	return
}

// Deletes

func (t *BTree) DeleteLeftSubtree() {
	t.left = nil
}

func (t *BTree) DeleteRightSubtree() {
	t.right = nil
}

func (t *BTree) DeleteLeftChild() {
	// TO DO
}

func (t *BTree) DeleteRightChild() {
	// TO DO
}

//Helper Functions

func (t BTree) IsLeaf() bool {
	return (t.left == nil && t.right == nil)
}

func (t BTree) HasLeftChild() bool {
	return t.left != nil
}

func (t BTree) HasRightChild() bool {
	return t.right != nil
}

func (t BTree) Depth() (d int64) {
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

func (t BTree) prefixDo(f func(interface{})) {
	f(t.val)
	if t.left != nil {
		t.left.prefixDo(f)
	}
	if t.right != nil {
		t.right.prefixDo(f)
	}
}

func (t BTree) inOrderDo(f func(interface{})) {
	if t.left != nil {
		t.left.inOrderDo(f)
	}
	f(t.val)
	if t.right != nil {
		t.right.inOrderDo(f)
	}
}

func (t BTree) postfixDo(f func(interface{})) {
	if t.left != nil {
		t.left.postfixDo(f)
	}
	if t.right != nil {
		t.right.postfixDo(f)
	}
	f(t.val)
}

func (t BTree) breadthFirstDo(f func(interface{})) {
	q := MakeQueue()
	q.Push(&t)

	for !q.Empty() {
		p := q.Front().(*BTree)
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

func (t BTree) depthFirstDo(f func(interface{})) {
	s := MakeStack()
	s.Push(&t)

	for !s.Empty() {
		p := s.Top().(*BTree)
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

func (t BTree) prefixPrint() {
	t.prefixDo(func(x interface{}) {
		fmt.Println(x, " ")
	})
	fmt.Println("")
}

func (t BTree) postfixPrint() {
	t.postfixDo(func(x interface{}) {
		fmt.Println(x, " ")
	})
	fmt.Println("")
}

func (t BTree) inOrderPrint() {
	t.inOrderDo(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println("")
}

func (t BTree) BFPrint() {
	t.breadthFirstDo(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println("")
}

func (t BTree) DFPrint() {
	t.depthFirstDo(func(x interface{}) {
		fmt.Print(x, " ")
	})
	fmt.Println("")
}

//Pretty Printer
//http://stackoverflow.com/questions/4965335/how-to-print-binary-tree-diagram

func (t BTree) PrettyPrint() {
	maxLevel := t.Depth()
	nodes := make([]*BTree, 0)
	nodes = append(nodes, &t)
	t.prettyPrintHelper(nodes, 1, maxLevel+1)
}

func allNull(nodes []*BTree) bool {
	for _, n := range nodes {
		if n != nil {
			return false
		}
	}
	return true
}

func printWhiteSpace(num int64) {
	for i := int64(0); i < num; i++ {
		fmt.Print(" ")
	}
}

func (t BTree) prettyPrintHelper(nodes []*BTree, level, maxLevel int64) {
	if len(nodes) == 0 || allNull(nodes) {
		return
	}

	floor := maxLevel - level
	edgeLines := PowInt(2, MaxInt(floor-1, 0))
	firstSpaces := PowInt(2, floor) - 1
	betweenSpaces := PowInt(2, floor+1) - 1

	printWhiteSpace(firstSpaces)

	newNodes := make([]*BTree, 0)
	for _, n := range nodes {
		xtra := int64(0)
		if n != nil {
			extra, _ := fmt.Print(n.val)
			xtra = int64(extra - 1)
			newNodes = append(newNodes, n.left)
			newNodes = append(newNodes, n.right)
		} else {
			newNodes = append(newNodes, nil, nil)
			fmt.Print(" ")
		}
		printWhiteSpace(betweenSpaces - xtra)
	}
	fmt.Println("")

	for i := int64(1); i <= edgeLines; i++ {
		for j := int64(0); j < int64(len(nodes)); j++ {
			printWhiteSpace(firstSpaces - i)
			if nodes[j] == nil {
				printWhiteSpace(edgeLines*2 + i + 1)
				continue
			}
			if nodes[j].left != nil {
				fmt.Print("/")
			} else {
				fmt.Print(" ")
			}
			printWhiteSpace(2*i - 1)
			if nodes[j].right != nil {
				fmt.Print("\\")
			} else {
				fmt.Print(" ")
			}
			printWhiteSpace(edgeLines*2 - i)

		}
		fmt.Println("")

	}
	t.prettyPrintHelper(newNodes, level+1, maxLevel)

}
