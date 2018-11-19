package title207

// 检测是否有环
// 拓扑排序
// dfs访问各个节点
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 状态 2 表示正在访问的编号，也就是出度暂时不为0的编号
	// 状态 1 表示已经访问完毕的编号，也就是出度已经为0的编号
	// 状态 0 表示节点还没有进行访问的

	pre := [][]int{}
	for i := 0; i < numCourses; i++ {
		p := make([]int, numCourses)
		pre = append(pre, p)
	}

	for _, p := range prerequisites {
		pre[p[1]][p[0]] = 1
	}

	condition := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		condition[i] = 0
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, pre, condition) {
			return false
		}
	}
	return true
}

func dfs(cur int, prerequisites [][]int, condition map[int]int) bool {

	if condition[cur] == 2 {
		return false
	}

	if condition[cur] == 1 {
		return true
	}

	// 1、标记为正在遍历
	condition[cur] = 2

	// 2、遍历下一个节点
	for i := 0; i < len(prerequisites); i++ {
		if prerequisites[cur][i] == 1 {
			if !dfs(i, prerequisites, condition) {
				return false
			}
		}
	}

	condition[cur] = 1
	return true
}
