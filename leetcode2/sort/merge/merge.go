package main

import (
	"fmt"
)

func main() {
	test := []int{4, 3, 5, 2, 6, 3, 5, 3, 5, 3}
	Sort(test, 0, len(test)-1)
	fmt.Println(test)
}

func Sort(numbers []int, head int, tail int) {
	if head < tail {
		middle := (head + tail) / 2
		Sort(numbers, head, middle)
		Sort(numbers, middle+1, tail)
		Merge(numbers, head, middle, tail)
	}
}

// 归并算法
func Merge(numbers []int, head int, middle int, tail int) {

	// 处理单个元素
	if middle == head {
		if numbers[tail] < numbers[head] {
			tmp := numbers[tail]
			numbers[tail] = numbers[head]
			numbers[head] = tmp
		}
		return
	}

	// 归并两个有序数列
	m := middle + 1
	first := head
	var tmp []int
	for first <= middle && m <= tail {
		if numbers[first] < numbers[m] {
			tmp = append(tmp, numbers[first])
			first++
		} else if numbers[m] < numbers[first] {
			tmp = append(tmp, numbers[m])
			m++
		} else {
			tmp = append(tmp, numbers[first])
			tmp = append(tmp, numbers[m])
			m++
			first++
		}
	}

	if m <= tail {
		for i := m; i <= tail; i++ {
			tmp = append(tmp, numbers[i])
		}
	} else if first <= middle {
		for i := first; i <= middle; i++ {
			tmp = append(tmp, numbers[i])

		}
	}
	for i := head; i <= tail; i++ {
		numbers[i] = tmp[i-head]
	}
}
