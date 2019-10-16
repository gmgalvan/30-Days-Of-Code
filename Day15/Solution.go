package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Data int
	next *Node
}

func InsertNode(head *Node, node *Node, data int) (*Node, *Node) {
	var newNode Node
	var next *Node
	if head == nil {
		newNode.Data = data
		head = &newNode
		next = &newNode

	} else {
		newNode.Data = data
		node.next = &newNode
		next = &newNode
	}

	return head, next
}

func Solution(head *Node) {
	fmt.Print(head.Data, " ")
	next := head.next
	for next != nil {
		print(next.Data, " ")
		next = next.next
	}
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	var node *Node
	var head *Node
	for i := 0; i < n; i++ {
		in.Scan()
		input := in.Text()
		data, _ := strconv.Atoi(input)
		head, node = InsertNode(head, node, data)
	}
	Solution(head)
}
