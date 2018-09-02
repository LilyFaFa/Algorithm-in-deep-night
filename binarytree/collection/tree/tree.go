package tree

type Node struct {
	Item      int
	LeftNode  *Node
	RightNode *Node
}

func NewTree() *Node {
	test := &Node{
		Item: 1,
		LeftNode: &Node{
			Item: 2,
			LeftNode: &Node{
				Item: 3,
				LeftNode: &Node{
					Item: 4,
				},
			},
			RightNode: &Node{
				Item: 5,
				LeftNode: &Node{
					Item: 34,
				},
			},
		},
		RightNode: &Node{
			Item: 5,
			RightNode: &Node{
				Item: 6,
				RightNode: &Node{
					Item: 7,
				},
			},
			LeftNode: &Node{
				Item: 23,
				RightNode: &Node{
					Item: 34,
					LeftNode: &Node{
						Item: 25,
					},
				},
				LeftNode: &Node{
					Item: 35,
				},
			},
		},
	}
	return test
}

func NewSearchTree() *Node {
	test := &Node{
		Item: 10,
		LeftNode: &Node{
			Item: 6,
			LeftNode: &Node{
				Item: 4,
			},
			RightNode: &Node{
				Item: 8,
			},
		},
		RightNode: &Node{
			Item: 14,
			LeftNode: &Node{
				Item: 12,
			},
			RightNode: &Node{
				Item: 16,
			},
		},
	}
	return test
}
