package depthfirst

import (
	"github.com/LilyFaFa/Algorithm-in-deep-night/collections"
)

// GraphSearch
// 从图的某个节点出发，访问连通图的所有节点，所有节点只被访问一次，使用深度优先算法
// 默认图的节点是从0开始的连续的正整数
// 输入顶点邻接矩阵，输出访问顺序的队列

// 访问某个节点
func AccessNode(vv [][]int, currentNode int, q *collections.Stack, visited []bool) {
	result := collections.Queue{}
	result.Enqueue(currentNode)
	for i, j := range vv[currentNode] {
		if j == 1 && !visited[i] {
			//进队列的设置为已经访问，因为，不设置可能会反复进队列
			visited[i] = true
			q.Push(i)
		}
	}
}

// 深度优先搜索
func DepthFirst(vv [][]int) *collections.Queue {

	// 结果队列
	result := &collections.Queue{}

	// 中间访问栈
	q := &collections.Stack{}

	// 初始节点设置为0
	currentNode := 0

	// 记录节点是否被访问过
	visited := []bool{}
	for i := 0; i < len(vv); i++ {
		visited = append(visited, false)
	}

	// 访问起始节点
	result.Enqueue(currentNode)
	visited[0] = true
	AccessNode(vv, currentNode, q, visited)

	for !q.IsEmpty() {
		current, _ := q.Pop()
		// 断言
		currentNode = current.(int)
		result.Enqueue(currentNode)
		visited[currentNode] = true
		AccessNode(vv, currentNode, q, visited)
	}

	return result
}
