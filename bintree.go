package ds

type BinaryTree struct {
	left  *BinaryTree
	right *BinaryTree
	val   interface{}
}

func MakeBinaryTree(val interface{}) *BinaryTree {
	var b BinaryTree
	return &b
}

func MakeBinaryTreeWithLeftSubtree(valToAdd interface{}, lefty *BinaryTree) *BinaryTree {
	var b BinaryTree
	b.val = valToAdd
	b.left = lefty
	return &b
}

func MakeBinaryTreeWithRightSubtree(valToAdd interface{}, righty *BinaryTree) *BinaryTree {
	var b BinaryTree
	b.val = valToAdd
	b.right = righty
	return &b
}

func MakeBinaryTreeWithSubtrees(valToAdd interface{}, lefty, righty *BinaryTree) *BinaryTree {
	var b BinaryTree
	b.val = valToAdd
	b.left = lefty
	b.right = righty
	return &b
}

func (b *BinaryTree) IsLeaf() bool {
	return (b.left == nil && b.right == nil)
}

func (b BinaryTree) HasLeftChild() bool {
	return b.left != nil
}

func (b BinaryTree) HasRightChild() bool {
	return b.right != nil
}

func (b BinaryTree) Depth() int {
	var d int
	if b.IsLeaf() {
		d = 1
	} else {
		if b.left == nil {
			d = (*b.right).Depth() + 1
		} else if b.right == nil {
			d = (*b.left).Depth() + 1
		} else {
			ld := b.left.Depth() + 1
			rd := b.right.Depth() + 1

			if ld > rd {
				d = ld
			} else {
				d = rd
			}
		}
	}

	return d

}

func (b *BinaryTree) RemoveLeftSubtree() {
	b.left = nil
}

func (b *BinaryTree) RemoveRightSubtree() {
	b.right = nil
}

//Dos

func (t *BinaryTree) prefixDo(f func(interface{})) {
	f(t.val)
	if t.left != nil {
		t.left.prefixDo(f)
	}
	if t.right != nil {
		t.right.prefixDo(f)
	}
}

func (t *BinaryTree) inOrderDo(f func(interface{})) {
	if t.left != nil {
		t.left.inOrderDo(f)
	}
	f(t.val)
	if t.right != nil {
		t.right.inOrderDo(f)
	}
}

func (t *BinaryTree) postfixDo(f func(interface{})) {
	if t.left != nil {
		t.left.postfixDo(f)
	}
	if t.right != nil {
		t.right.postfixDo(f)
	}
	f(t.val)
}

func (t *BinaryTree) breadthFirstDo(f func(interface{})) {
	q := MakeQueue()
	q.Push(&t)

	for !q.Empty() {
		p := q.Front().(*BinaryTree)
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

func (t *BinaryTree) depthFirstDo(f func(interface{})) {
	s := MakeStack()
	s.Push(&t)

	for !s.Empty() {
		p := s.Top().(*BinaryTree)
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
