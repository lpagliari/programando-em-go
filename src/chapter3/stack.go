package main

import (
	"errors"
	"fmt"
)

func main() {
	stack := Stack{}

	fmt.Println("Stack created. Size:", stack.Size())
	fmt.Println("Stack is empty?", stack.Empty())

	stack.Push("Go")
	stack.Push(2009)
	stack.Push(3.14)
	stack.Push("End")

	fmt.Println("Stack changed. Size:", stack.Size())
	fmt.Println("Stack is empty?", stack.Empty())

	for !stack.Empty() {
		value, _ := stack.Pop()

		fmt.Println("Value popped:", value)
		fmt.Println("Size:", stack.Size())
		fmt.Println("Stack is empty?", stack.Empty())
	}

	// force the error
	_, err := stack.Pop()
	if err != nil {
		fmt.Println("Error found: ", err)
	}
}

type Stack struct {
	values []interface{}
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) Empty() bool {
	return stack.Size() == 0
}

func (stack *Stack) Push(value interface{}) {
	stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
	if stack.Empty() {
		return nil, errors.New("Empty stack!")
	}

	lastIndex := stack.Size() - 1
	value := stack.values[lastIndex]
	stack.values = stack.values[:lastIndex]
	return value, nil
}
