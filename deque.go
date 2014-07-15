package ds

import (
	"fmt"
)

type Deque struct {
	data []interface{}
}

func MakeDeque() *Deque {
	var d Deque
	d.Init()
	return &d
}

func (d *Deque) Init() {
	d.data = make([]interface{}, 0)
}

func (d Deque) Empty() bool {
	return len(d.data) == 0
}

func (d Deque) Size() int {
	return len(d.data)
}

func (d Deque) Front() interface{} {
	if d.Empty() {
		return 0
	}
	return d.data[0]
}

func (d Deque) Back() interface{} {
	return d.data[len(d.data)-1]
}

func (d *Deque) Push_Front(toAdd interface{}) {
	newElem := make([]interface{}, 0)
	newElem = append(newElem, toAdd)
	for x := 0; x < len(d.data); x++ {
		newElem = append(newElem, d.data[x])
	}
	d.data = newElem
}

func (d *Deque) Push_Back(toAdd interface{}) {
	d.data = append(d.data, toAdd)
}

func (d *Deque) Pop_Front() interface{} {
	toReturn := d.data[0]
	d.data = d.data[1:]
	return toReturn
}

func (d *Deque) Pop_Back() interface{} {
	toReturn := d.data[len(d.data)-1]
	d.data = d.data[0 : len(d.data)-1]
	return toReturn
}

func (d Deque) Print() {
	fmt.Println("Deque of size: ", len(d.data))
	for x := 0; x < len(d.data); x++ {
		fmt.Print(d.data[x], ", ")
	}
	fmt.Println("")

}
