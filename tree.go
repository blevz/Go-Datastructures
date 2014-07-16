package ds

import (
	_ "fmt"
)

type Tree struct {
	children []*Tree
	val      interface{}
}

func MakeTree(valToAdd interface{}) *Tree {
	var t Tree
	t.val = valToAdd
	t.children = make([]*Tree, 0)
	return &t
}

func (t Tree) IsLeaf() bool {
	return len(t.children) == 0
}

func (t *Tree) AddChild(child *Tree) {
	t.children = append(t.children, child)
}

func (t *Tree) AddChildWithValue(valToAdd interface{}) {
	child := MakeTree(valToAdd)
	t.AddChild(child)
}

func (t *Tree) GetChildren() []*Tree {
	return t.children
}

/*
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

}*/
