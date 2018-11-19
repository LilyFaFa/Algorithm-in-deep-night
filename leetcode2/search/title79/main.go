package search

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			finded := false
			search(board, i, j, word, "", &finded)
			if finded {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i int, j int, target string, cur string, searched *bool) {
	// 1.1 递归退出条件：找到了字符串
	if *searched {
		return
	}
	// 1.2 递归退出条件：i和j超出边界
	if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 {
		return
	}
	// 2 处理递归
	if board[i][j] == target[len(cur)] {
		if len(target) == len(cur)+1 {
			*searched = true
			return
		}
		// 访问过的board[i][j]，我们置成0
		tmp := board[i][j]
		board[i][j] = '0'
		search(board, i+1, j, target, cur+string(target[len(cur)]), searched)
		search(board, i, j+1, target, cur+string(target[len(cur)]), searched)
		search(board, i-1, j, target, cur+string(target[len(cur)]), searched)
		search(board, i, j-1, target, cur+string(target[len(cur)]), searched)
		// 复原
		board[i][j] = tmp
	}

}
