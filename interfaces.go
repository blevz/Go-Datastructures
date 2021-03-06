package ds

type Set_I interface {
	Size() int
	Empty() bool
	AddElem(interface{})
	DelElem(interface{})
	Exist(interface{}) bool
}

type Multiset_I interface {
	Set_I
	Decrement_Element(interface{})
	Size_With_Duplicates() int
	Count(interface{}) int
}

type Extended_Set_I interface {
	Set_I
	Union(Extended_Set_I) Extended_Set_I
	Intersect(Extended_Set_I) Extended_Set_I
	Difference(Extended_Set_I) Extended_Set_I
}

type Disjoint_Set_I interface {
	Find(int) int
	Union(int, int)
	Print()
}

type Priority_Queue_I interface {
	Push(interface{}, float64)
	Pop() interface{}
}

type Deque_I interface {
	Empty() bool
	Size() int
	Front() interface{}
	Back() interface{}
	Push_Front(interface{})
	Push_Back(interface{})
	Pop_Front() interface{}
	Pop_Back() interface{}
	Print()
}

type Queue_I interface {
	Empty() bool
	Size() int
	Front() interface{}
	Push(interface{})
	Pop() interface{}
	Print()
}

type Stack_I interface {
	Empty() bool
	Size() int
	Top() interface{}
	Push(interface{})
	Pop() interface{}
}

type Two_Dimensional_Array_I interface {
	Size() (int, int)
	SetElem(int, int, interface{})
	GetElem(int, int) interface{}
}

type Tree_I interface {
	IsLeaf() bool
	AddChild(Tree_I)
	AddChildWithValue(interface{})
	GetChildren() []Tree_I
}

type Binary_Tree_I interface {
	GetLeftSubtree() Binary_Tree_I
	GetRightSubtree() Binary_Tree_I
}

type Graph_I interface {
}
