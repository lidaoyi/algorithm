package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

func main() {
	value := 1
	obj := Constructor()
	obj.AppendTail(value)
	param_2 := obj.DeleteHead()
	fmt.Println(param_2)
}

type CQueue struct {
	s1 *stack.Stack
	s2 *stack.Stack
}

func Constructor() CQueue {
	return CQueue{stack.New(), stack.New()}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.s2.Len() > 0 {
		return this.s2.Pop().(int)
	} else {
		return -1
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
