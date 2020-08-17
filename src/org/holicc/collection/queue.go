package main

import (
	"fmt"
)

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */

func main() {
	queue := Constructor()

	queue.Push_back(46)
	fmt.Println(queue.Pop_front())
	queue.Push_back(46)
	queue.Push_back(46)
	queue.Push_back(46)
	queue.Push_back(46)
	queue.Push_back(46)
	fmt.Println(queue.Pop_front())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
	fmt.Println(queue.Max_value())
}

type MaxQueue struct {
	que    Queue
	maxque Queue
}

type Queue []int

func Constructor() MaxQueue {
	return MaxQueue{
		que:    make(Queue, 0),
		maxque: make(Queue, 0),
	}
}

func (q Queue) Back() int {
	return q[len(q)-1]
}

func (this *MaxQueue) Max_value() int {
	if len(this.que) == 0 {
		return -1
	}
	return this.maxque[0]
}

func (this *MaxQueue) Push_back(value int) {
	for len(this.maxque) != 0 && value > this.maxque.Back() {
		this.maxque = this.maxque[:len(this.maxque)-1]
	}
	this.maxque = append(this.maxque, value)
	this.que = append(this.que, value)
}

func (this *MaxQueue) Pop_front() int {
	key := -1
	if len(this.que) != 0 {
		key = this.que[0]
		this.que = this.que[1:]
		if key == this.maxque[0] {
			this.maxque = this.maxque[1:]
		}
	}
	return key
}
