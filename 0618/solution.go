package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "1-2--3---4-5--6---7"

	arr := strings.Split(str, "-")
	for _, x := range arr {
		fmt.Printf("x: %s \n", x)
	}

	fmt.Println(arr)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//"1-2--3--4-5--6--7"
func recoverFromPreorder(S string) *TreeNode {
	node := &TreeNode{}
	stack := NewStack()

	flagCount := 0
	temp := ""

	arr := strings.Split(S, "-")


	return node
}

func parseInt(s string) int {
	v, _ := strconv.ParseInt(s, 10, 32)
	return int(v)
}

type Stack struct {
	arr []*TreeNode
	top int
}

func NewStack() *Stack {
	return &Stack{make([]*TreeNode, 64), 0}
}

func (self Stack) put(item *TreeNode) {
	self.arr = append(self.arr, item)
	self.top = self.top + 1
}

func (self Stack) pop() *TreeNode {
	item := self.arr[self.top-1]
	self.top = self.top - 1
	return item
}
