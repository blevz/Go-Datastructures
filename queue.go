package ds

import "fmt"

type Queue struct {
	data []interface{}
}

func MakeQueue() *Queue {
	var q Queue
	q.Init()
	return &q
}

func (q Queue) Init() {
	q.data = make([]interface{}, 0)
}

func (q Queue) Front() interface{} {
	if q.Empty() {
		return 0
	}
	return q.data[0]
}

func (q Queue) Size() int {
	return len(q.data)
}

func (q Queue) Empty() bool {
	return len(q.data) == 0
}

func (q *Queue) Pop() interface{} {
	if q.Empty() {
		return 0
	}
	toReturn := q.data[0]
	q.data = q.data[1:q.Size()]
	return toReturn
}

func (q *Queue) Push(x interface{}) {
	q.data = append(q.data, x)
}

func (q Queue) Print() {
	for _, x := range q.data {
		fmt.Print(x, " ")
	}
	fmt.Println("")
}
