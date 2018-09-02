package main

import (
	"fmt"

	. "github.com/LilyFaFa/Algorithm-in-deep-night/collection/binarytree/tree"
)

func convert(node *Node, linkHead **Node, linkTail **Node) {
	if node == nil {
		return
	}
	convert(node.LeftNode, linkHead, linkTail)
	node.LeftNode = *linkTail
	if *linkTail != nil {
		(*linkTail).RightNode = node
	}
	if *linkHead == nil {
		*linkHead = node
	}
	*linkTail = node

	convert(node.RightNode, linkHead, linkTail)
}

func main() {

	n := NewSearchTree()
	fmt.Println("covert......")
	var first, last **Node
	first = new(*Node)
	last = new(*Node)
	fmt.Println(*first, *last)
	convert(n, first, last)
	f := *first

	for {
		if f != nil {
			fmt.Println(f.Item)
			f = f.RightNode
		} else {
			break
		}
	}
}
