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
func MakeTreeWithSubtrees(lefty *Tree, righty *Tree, val interface{}) (t Tree) {
	t.right = righty
	t.left = lefty
	t.val = val
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
		t.left.inOrderDo(f)
	}
	f(t.val)
	if t.right != nil {
		t.right.inOrderDo(f)
	}
}

func (t Tree) postfixDo(f func(interface{})) {
	if t.left != nil {
		t.left.postfixDo(f)
	}
	if t.right != nil {
		t.right.postfixDo(f)
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

//Pretty Printer
//http://stackoverflow.com/questions/4965335/how-to-print-binary-tree-diagram

func (t Tree) PrettyPrint() {
	maxLevel := t.Depth()
	nodes := make([]*Tree, 0)
	nodes = append(nodes, &t)
	t.prettyPrintHelper(nodes, 1, maxLevel+1)
}

func allNull(nodes []*Tree) bool {
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

func (t Tree) prettyPrintHelper(nodes []*Tree, level, maxLevel int64) {
	if len(nodes) == 0 || allNull(nodes) {
		return
	}

	floor := maxLevel - level
	edgeLines := PowInt(2, MaxInt(floor-1, 0))
	firstSpaces := PowInt(2, floor) - 1
	betweenSpaces := PowInt(2, floor+1) - 1

	printWhiteSpace(firstSpaces)

	newNodes := make([]*Tree, 0)
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
