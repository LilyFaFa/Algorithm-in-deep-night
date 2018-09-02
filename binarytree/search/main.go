package main

import (
	"fmt"

	. "github.com/LilyFaFa/Algorithm-in-deep-night/binarytree/collection/queue"
	. "github.com/LilyFaFa/Algorithm-in-deep-night/binarytree/collection/tree"
)

func firstSearch(root *Node) {
	if root != nil {
		fmt.Println(root.Item)
		firstSearch(root.LeftNode)
		firstSearch(root.RightNode)
	}
}

func middleSearch(root *Node) {
	if root != nil {
		middleSearch(root.LeftNode)
		fmt.Println(root.Item)
		middleSearch(root.RightNode)
	}
}

func lastSearch(root *Node) {
	if root != nil {
		lastSearch(root.LeftNode)
		lastSearch(root.RightNode)
		fmt.Println(root.Item)
	}
}

func leavelSearch(root *Node, queue *Queue) {
	if root != nil {
		fmt.Println(root.Item)
		if root.LeftNode != nil {
			queue.EnQueue(root.LeftNode)
		}
		if root.RightNode != nil {
			queue.EnQueue(root.RightNode)
		}

		node, _ := queue.DeQueue()
		value, _ := node.(*Node)
		leavelSearch(value, queue)
	}
}

func main() {
	n := NewTree()
	fmt.Println("middleSearch......")
	middleSearch(n)
	fmt.Println("firstSearch......")
	firstSearch(n)
	fmt.Println("lastSearch......")
	lastSearch(n)
	fmt.Println("leavelSearch.....")
	q := NewQueue()
	leavelSearch(n, q)
}
