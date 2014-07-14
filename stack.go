package ds

import (
	"fmt"
)

type Stack struct {
	data []interface{}
}

func MakeStack() *Stack {
	var s Stack
	s.Init()
	return &s
}

func (s *Stack) Init() {
	s.data = make([]interface{}, 0)
}

func (s Stack) Top() interface{} {
	if s.Empty() {
		return 0
	}
	return s.data[s.Size()-1]
}

func (s Stack) Size() int {
	return len(s.data)
}

func (s Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return 0
	}
	toReturn := s.data[s.Size()-1]
	s.data = s.data[0 : s.Size()-1]
	return toReturn
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s Stack) Print() {
	for _, x := range s.data {
		fmt.Print(x, " ")
	}
	fmt.Println("")
}
