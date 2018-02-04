package lcs

//http://www.cnblogs.com/huangxincheng/archive/2012/11/11/2764625.html
func LCS(x string, y string) (string, int) {
	// 记录自己的上一个点的来源，0是指左上方，1是指上方，2是指左方
	// 左上方的也就是标示该点是要相同的点
	var note [][]int

	// 记录中间过程中的中间结果
	var mem [][]int

	// x和y长度为零,没有公共子字符串
	if len(x) == 0 || len(y) == 0 {
		return "", 0
	}

	// 初始化note和mem
	if x[0] == y[0] {
		notetemp := []int{0}
		memtemp := []int{1}
		note = append(note, notetemp)
		mem = append(mem, memtemp)
	} else {
		notetemp := []int{1}
		memtemp := []int{0}
		note = append(note, notetemp)
		mem = append(mem, memtemp)
	}

	for i, v1 := range x {
		for j, v2 := range y {
			if i == 0 && j == 0 {
				continue
			}
			if i != 0 && j != 0 {
				if v1 == v2 {
					mem[i] = append(mem[i], mem[i-1][j-1]+1)
					note[i] = append(note[i], 0)
				} else {
					if mem[i-1][j] > mem[i][j-1] {
						mem[i] = append(mem[i], mem[i-1][j])
						note[i] = append(note[i], 1)
					} else {
						mem[i] = append(mem[i], mem[i][j-1])
						note[i] = append(note[i], 2)
					}
				}

			} else if j == 0 {
				memtemp := []int{}
				notetemp := []int{}
				if v1 == v2 {
					memtemp = append(memtemp, 1)
					notetemp = append(notetemp, 0)
				} else {
					memtemp = append(memtemp, 0)
					notetemp = append(notetemp, 1)
				}
				mem = append(mem, memtemp)
				note = append(note, notetemp)

			} else if i == 0 {
				if v1 == v2 {
					mem[0] = append(mem[0], 1)
					note[0] = append(note[0], 0)
				} else {
					mem[0] = append(mem[0], 0)
					note[0] = append(note[0], 1)
				}
			}
		}
	}
	z := ""
	SubSequence(note, x, len(x)-1, len(y)-1, &z)
	return z, len(z)
}

func SubSequence(note [][]int, x string, a int, b int, z *string) {
	if a == 0 || b == 0 {
		if note[a][b] == 0 {
			*z = string(x[a]) + *z
		}
		return
	} else {
		n := note[a][b]
		if n == 0 {
			*z = string(x[a]) + *z
			SubSequence(note, x, a-1, b-1, z)
		} else if n == 1 {
			SubSequence(note, x, a-1, b, z)
		} else {
			SubSequence(note, x, a, b-1, z)
		}
	}

}
