package title841

// 看是否有多于1个连同子图，使用dfs，一次搜索之后，如果还是有
// 没有访问过的节点，那么认为有多于一个连同图

func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]int)
	for i := 0; i < len(rooms); i++ {
		visited[i] = 0
	}
	for i := 0; i < len(rooms); i++ {

	}
}

func dfs(grape [][]int, visited map[int]int) {

}
