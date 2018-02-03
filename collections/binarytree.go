package collections

import (
	"errors"
)

type Node struct {
	Item       interface{}
	ParentNode *Node
	LeftNode   *Node
	RightNode  *Node
}

// To define a tree, We use a topNode and a compare function
type BinaryTree struct {
	TopNode *Node
	Compare func(a, b interface{}) bool
}

func NewBinaryTree(topItem interface{}, compare func(a, b interface{}) bool) *BinaryTree {
	TopNode := &Node{
		Item: topItem,
	}
	return &BinaryTree{
		TopNode: TopNode,
		Compare: compare,
	}
}

func (b *BinaryTree) InsertNode(n *Node, item interface{}) {
	if b.Compare(n.LeftNode, item) {
		if n.LeftNode == nil {
			n.LeftNode = &Node{
				Item:       item,
				ParentNode: n,
			}
			return
		} else {
			b.InsertNode(n.LeftNode, item)
		}
	} else {
		if n.RightNode == nil {
			n.RightNode = &Node{
				Item:       item,
				ParentNode: n,
			}
			return
		} else {
			b.InsertNode(n.RightNode, item)
		}
	}

}

func (b *BinaryTree) Insert(item interface{}) error {
	if b.TopNode != nil {
		errors.New("current binarytree is empty")
	}
	b.InsertNode(b.TopNode, item)
	return nil
}

func (b *BinaryTree) Remove(node *Node) error {
	if node.LeftNode == nil && node.RightNode == nil {
		if node.ParentNode != nil {
			if node.ParentNode.LeftNode == node {
				node.ParentNode.LeftNode = nil
			} else {
				node.ParentNode.RightNode = nil
			}
			return nil
		} else {
			node = nil
			return errors.New("delete the top node,current tree is empty")
		}
	}

	if node.LeftNode != nil && node.RightNode == nil {
		if node.ParentNode != nil {
			if node.ParentNode.LeftNode == node {
				node.ParentNode.LeftNode = node.LeftNode
			} else {
				node.ParentNode.RightNode = node.LeftNode
			}
		} else {
			b.TopNode = node.LeftNode
			node.LeftNode.ParentNode = nil
		}
		return nil
	}

	if node.LeftNode == nil && node.RightNode != nil {
		if node.ParentNode != nil {
			if node.ParentNode.LeftNode == node {
				node.ParentNode.LeftNode = node.RightNode
			} else {
				node.ParentNode.RightNode = node.RightNode
			}
		} else {
			b.TopNode = node.RightNode
			node.RightNode.ParentNode = nil
		}
		return nil
	}

	selectedNode := findNodeWithoutLeftNode(node.RightNode)
	selectedNode.LeftNode = node.LeftNode

	if selectedNode.ParentNode == node {
		if node.ParentNode != nil {
			selectedNode.LeftNode = node.LeftNode
			selectedNode.ParentNode = node.ParentNode
			if node.ParentNode.LeftNode == node {
				node.ParentNode.LeftNode = selectedNode
			} else {
				node.ParentNode.RightNode = selectedNode
			}
		} else {
			selectedNode.ParentNode = nil
			b.TopNode = selectedNode
		}

	} else {
		if selectedNode.RightNode != nil {
			selectedNode.RightNode.ParentNode = selectedNode.ParentNode
		}
		selectedNode.ParentNode.LeftNode = selectedNode.RightNode

		node.LeftNode.ParentNode = selectedNode
		node.RightNode.ParentNode = selectedNode
		selectedNode.LeftNode = node.LeftNode
		selectedNode.RightNode = node.RightNode

		if node.ParentNode != nil {
			selectedNode.ParentNode = node.ParentNode
		} else {
			selectedNode.ParentNode = nil
			b.TopNode = selectedNode
		}
	}
	return nil

}

func findNodeWithoutLeftNode(node *Node) *Node {
	if node.LeftNode == nil {
		return node
	} else {
		findNodeWithoutLeftNode(node.LeftNode)
	}
	return nil
}
