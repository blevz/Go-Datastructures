package ds

type Queue struct {
	data []interface{}
}

func MakeQueue() (q Queue) {
	q.Init()
	return
}

func (q Queue) Init() {
	q.data = make([]interface{}, 0)
}

func (q Queue) Front() interface{} {
	return q.data[0]
}

func (q Queue) Size() int {
	return len(q.data)
}

func (q *Queue) Pop() {
	q.data = q.data[1:q.Size()]
}

func (q *Queue) Push(x interface{}) {
	q.data = append(q.data, x)
}
