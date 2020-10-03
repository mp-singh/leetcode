package main

import "log"

type MinStack struct {
	Stack    []int
	StackMin []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	m.Stack = append(m.Stack, x)
	if len(m.StackMin) == 0 || x <= m.StackMin[len(m.StackMin)-1] {
		m.StackMin = append(m.StackMin, x)
	}
}

func (m *MinStack) Pop() {
	x := m.Stack[len(m.Stack)-1]
	m.Stack = m.Stack[:len(m.Stack)-1]
	if x <= m.StackMin[len(m.StackMin)-1] {
		_, m.StackMin = m.StackMin[len(m.StackMin)-1], m.StackMin[:len(m.StackMin)-1]
	}
}

func (m *MinStack) Top() int {
	return m.Stack[len(m.Stack)-1]
}

func (m *MinStack) GetMin() int {
	if len(m.StackMin) == 0 {
		return 0
	}
	return m.StackMin[len(m.StackMin)-1]
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	log.Println("obj:", obj)
	log.Print("min:", obj.GetMin())
	log.Println("obj:", obj)
	obj.Pop()
	log.Print("min:", obj.GetMin())
	log.Println("obj:", obj)
	obj.Pop()
	log.Print("min:", obj.GetMin())
	log.Println("obj:", obj)
	obj.Pop()
	obj.GetMin()
	log.Println("obj:", obj)
	log.Println("obj:", obj)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMax();
 */
