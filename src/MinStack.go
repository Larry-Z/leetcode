package main

import (
	"container/list"
	"fmt"
)

type MinStack struct {
	data *list.List
}

type Node struct {
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{list.New()}
}

func (this *MinStack) Push(x int) {
	if this.data.Len() == 0 {
		this.data.PushFront(&Node{val: x, min: x})
	} else {
		front := this.data.Front().Value.(*Node)
		n := &Node{x, front.min}
		if x < front.min {
			n.min = x
		}
		this.data.PushFront(n)
	}
}

func (this *MinStack) Pop() {
	if this.data.Len() == 0 {
		return
	}
	this.data.Remove(this.data.Front())
}

func (this *MinStack) Top() int {
	if this.data.Len() == 0 {
		return -1
	} else {
		return this.data.Front().Value.(*Node).val
	}
}

func (this *MinStack) GetMin() int {
	if this.data.Len() == 0 {
		return -1
	} else {
		return this.data.Front().Value.(*Node).min
	}
}

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(-5)
	s.Push(9)
	s.Push(-10)
	fmt.Println(s.GetMin() == -10)
	s.Pop()
	fmt.Println(s.GetMin() == -5)
	fmt.Println(s.Top() == 9)
}
