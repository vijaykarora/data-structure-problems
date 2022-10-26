package main

import "math"

type MinStack struct {
	Items []int
}

func Constructor() MinStack {
	return MinStack{
		Items: make([]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	m.Items = append(m.Items, val)
}

func (m *MinStack) Pop() {
	index := len(m.Items) - 1
	m.Items = m.Items[:index]
}

func (m *MinStack) Top() int {
	index := len(m.Items) - 1
	return m.Items[index]
}

func (m *MinStack) GetMin() int {
	min := math.MaxInt
	for _, value := range m.Items {
		if value < min {
			min = value
		}
	}
	return min
}
