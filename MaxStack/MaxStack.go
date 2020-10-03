package main

import "log"

type MaxStack struct {
	Stack    []int
	StackMax []int
}

func Constructor() MaxStack {
	return MaxStack{}
}

func (m *MaxStack) Push(x int) {
	m.Stack = append(m.Stack, x)
	if len(m.StackMax) == 0 || x >= m.StackMax[len(m.StackMax)-1] {
		m.StackMax = append(m.StackMax, x)
	}
}

func (m *MaxStack) Pop() {
	x := m.Stack[len(m.Stack)-1]
	m.Stack = m.Stack[:len(m.Stack)-1]
	if x >= m.StackMax[len(m.StackMax)-1] {
		_, m.StackMax = m.StackMax[len(m.StackMax)-1], m.StackMax[:len(m.StackMax)-1]
	}
}

func (m *MaxStack) Top() int {
	return m.Stack[len(m.Stack)-1]
}

func (m *MaxStack) GetMax() int {
	if len(m.StackMax) == 0 {
		return 0
	}
	return m.StackMax[len(m.StackMax)-1]
}

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Push(0)
	obj.Push(3)
	obj.Push(0)
	log.Println("obj:", obj)
	log.Print("max:", obj.GetMax())
	log.Println("obj:", obj)
	obj.Pop()
	log.Print("max:", obj.GetMax())
	log.Println("obj:", obj)
	obj.Pop()
	log.Print("max:", obj.GetMax())
	log.Println("obj:", obj)
	obj.Pop()
	obj.GetMax()
	log.Println("obj:", obj)
	log.Println("obj:", obj)
}

/**
 * Your MaxStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMax();
 */
