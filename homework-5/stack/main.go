package main

import "fmt"

type Element struct {
	Data interface{}
	Neighbour *Element
}

type Stack struct {
	Last *Element
	depth int
}

func (stack *Stack) Push(data interface{}) {
	new := &Element{Data: data}
	if stack.depth > 0 { new.Neighbour = stack.Last }
	stack.Last = new
	stack.depth += 1
}

func (stack *Stack) Pop() (data interface{}) {
	if stack.depth == 0 { return nil }
	leaving := stack.Last
	stack.Last = leaving.Neighbour
	stack.depth -= 1
	leaving.Neighbour = nil
	return leaving.Data
}

func (stack *Stack) Depth() int {
	return stack.depth
}

func main () {
	st := &Stack{}
	st.Push("1st")
	st.Push("2nd")
	st.Push("3rd")

	fmt.Println("stack depth =", st.Depth())
	for {
		value := st.Pop()
		if value == nil { break }
		fmt.Println("got from stack:", value)
	}
}