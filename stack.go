package ds

type Stack struct {
	data []interface{}
}

func MakeStack() (s Stack) {
	s.Init()
	return
}

func (s *Stack) Init() {
	s.data = make([]interface{}, 0)
}

func (s Stack) Top() interface{} {
	return s.data[s.Size()-1]
}

func (s Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Pop() {
	s.data = s.data[0 : s.Size()-1]
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}