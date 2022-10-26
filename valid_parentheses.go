package main

import (
	"fmt"
	"github.com/vijaykarora/data-structure-problems/base"
)

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

func main() {
	s := "(])"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	stack := base.Stack{}
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		current := arr[i]

		if stack.IsEmpty() && (current == 41 || current == 125 || current == 93) {
			stack.Push(current)
			continue
		}

		if current == 40 || current == 123 || current == 91 {
			stack.Push(current)
			continue
		}

		top := stack.Top()
		if current == 41 && top == 40 {
			stack.Pop()
			continue
		}

		if current == 125 && top == 123 {
			stack.Pop()
			continue
		}

		if current == 93 && top == 91 {
			stack.Pop()
			continue
		}

		stack.Push(current)
	}
	return stack.IsEmpty()
}
