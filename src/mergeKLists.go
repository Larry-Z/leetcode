package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"time"
)

func mergeKLists(lists []*ListNode) *ListNode {
	start := new(ListNode)
	p := start
	for {
		min := math.MaxInt32
		var minNode *ListNode
		idx := 0
		for i, node := range lists {
			if node != nil && node.Val < min {
				min = node.Val
				minNode = node
				idx = i
			}
		}
		if minNode == nil {
			break
		}
		p.Next = minNode
		lists[idx] = minNode.Next
		p = p.Next
	}
	p.Next = nil
	return start.Next
}

func mergeKLists1(lists []*ListNode) *ListNode {
	start := new(ListNode)
	p := start
	m := make(map[int]*ListNode, len(lists))
	for i, node := range lists {
		m[i] = node
	}
	for {
		if len(m) == 0 {
			break
		}
		min := math.MaxInt32
		var minNode *ListNode
		idx := 0
		for i, node := range m {
			if node != nil && node.Val < min {
				min = node.Val
				minNode = node
				idx = i
			}
		}
		if minNode == nil {
			break
		}
		p.Next = minNode
		if minNode.Next == nil {
			delete(m, idx)
		} else {
			m[idx] = minNode.Next
		}
		p = p.Next
	}
	p.Next = nil
	return start.Next
}

func run(mergeFunc func([]*ListNode) *ListNode) {
	lists := []*ListNode{}
	row, col := 500, 1000
	for i := 0; i < row; i++ {
		row := make([]int, 0)
		for j := 0; j < col; j++ {
			row = append(row, rand.Intn(10000))
		}
		sort.Ints(row)
		//fmt.Println(row)
		lists = append(lists, sliceToList(row))
	}
	t1 := time.Now()
	head := mergeFunc(lists)
	//printLinkList(head)
	fmt.Println(head)
	t2 := time.Now()
	fmt.Println("method", runtime.FuncForPC(reflect.ValueOf(mergeFunc).Pointer()).Name(), "time consume:", t2.Sub(t1))

}
func rangeBenchmark() {
	m := make(map[string]string)
	l := make([]string, 0)
	size := 10000000
	for i := 0; i < size; i++ {
		str := strconv.Itoa(i)
		m[str] = str
		l = append(l, str)
	}
	t1 := time.Now()
	ret := make([]string, 0, size*2)
	for _, v := range m {
		ret = append(ret, v)
	}
	t2 := time.Now()
	fmt.Println("map time consume", t2.Sub(t1))
	t3 := time.Now()
	for _, v := range l {
		ret = append(ret, v)
	}
	t4 := time.Now()
	fmt.Println("slice time consume", t4.Sub(t3))
}

func main() {
	//run(mergeKLists)
	//run(mergeKLists1)
	//rangeBenchmark()
	mergeKLists([]*ListNode{nil, nil})
}
