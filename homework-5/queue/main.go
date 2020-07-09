package main

import "fmt"

type Element struct {
	Data interface{}
	Rear *Element
}

type Queue struct {
	First *Element
	Last *Element
	length int
}

func (queue *Queue) Add(data interface{}) {
	new := &Element{Data: data}
	if queue.length == 0 {
		queue.First = new
	} else {
		queue.Last.Rear = new
	}
	queue.Last = new
	queue.length += 1
}

func (queue *Queue) Get() (data interface{}) {
	if queue.length == 0 { return nil }
	leaving := queue.First
	queue.First = leaving.Rear
	queue.length -= 1
	leaving.Rear = nil
	return leaving.Data
}

func (queue *Queue) Length() int {
	return queue.length
}

func main () {
	qu := &Queue{}
	qu.Add("1st")
	qu.Add("2nd")
	qu.Add("3rd")

	fmt.Println("queue length =", qu.Length())
	for {
		value := qu.Get()
		if value == nil { break }
		fmt.Println("got from queue:", value)
	}
}