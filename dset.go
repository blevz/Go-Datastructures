package ds

// Simple union find data struction
// Utilizing an int splice to store
// the data
// Extendable and deletable, but usually the size is set at the start

import (
	"fmt"
)

type DisjointSet struct {
	data []int
}

func MakeDJSet(num int) *DisjointSet {
	var r DisjointSet
	r.data = make([]int, num)
	for x := 0; x < num; x++ {
		r.data[x] = x
	}
	return &r
}

func (s DisjointSet) Find(num int) int {
	return s.data[num]
}

func (s *DisjointSet) Union(a, b int) {
	replace := s.Find(a)
	with := s.Find(b)
	if replace != with {
		for val, _ := range s.data {
			if s.data[val] == replace {
				s.data[val] = with
			}
		}
	}
	return
}

func (s DisjointSet) Print() {
	fmt.Println("Disjoint Set of Size: ", len(s.data))
	for x := 0; x < len(s.data); x++ {
		fmt.Print(s.data[x], ", ")
	}
	fmt.Println("")
}
