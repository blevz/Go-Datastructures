package ds

import "fmt"

type CircleQueue struct {
	size         int
	startElement int
	data         []interface{}
}

func MakeCircleQueue() *CircleQueue {
	var cq CircleQueue
	cq.Init(1)
	cq.data = append(cq.data, 0)
	return &cq
}

func (cq CircleQueue) Init(initialCap int) {
	cq.data = make([]interface{}, initialCap)
	cq.size = 0
	cq.startElement = 0
}

func (cq CircleQueue) Size() int {
	return cq.size
}

func (cq CircleQueue) Front() interface{} {
	if cq.size == 0 {
		return 0
	}
	return cq.data[cq.startElement]
}

func (cq CircleQueue) Empty() bool {
	return cq.size == 0
}

func (cq *CircleQueue) Push(toAdd interface{}) {
	cq.data[(cq.startElement+cq.size)%len(cq.data)] = toAdd
	cq.size++
	if cq.size == len(cq.data) {
		for x := 0; x < cq.size; x++ {
			cq.data = append(cq.data, cq.data[cq.startElement])
			cq.startElement++
		}
	}
}

func (cq *CircleQueue) Pop() interface{} {
	if cq.Empty() {
		return 0
	}
	toReturn := cq.data[cq.startElement]
	cq.startElement++
	cq.size--
	if cq.startElement == len(cq.data) {
		cq.startElement = 0
	}
	return toReturn
}

func (cq CircleQueue) Print() {
	cq.PrintCircleQueue()
}

func (cq CircleQueue) PrintCircleQueue() {

	fmt.Println("Start: ", cq.startElement, " Size: ", cq.size)
	for x := 0; x < len(cq.data); x++ {
		fmt.Print(cq.data[x])
	}
	fmt.Println("")

}
