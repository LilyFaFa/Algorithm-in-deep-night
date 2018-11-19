package main

func numMatchingSubseq(S string, words []string) int {
	temps := make(map[int]int)
	for i := 0; i < len(words); i++ {
		temps[i] = 0
	}

	for i := 0; i < len(S); i++ {
		s := S[i]
		t := 0
		for i, w := range words {
			if temps[i] >= len(w) {
				t++
				continue
			} else {
				if w[temps[i]] == s {
					temps[i] = temps[i] + 1
				}
			}
		}
		if t >= len(words) {
			break
		}
	}
	result := 0
	for i, w := range words {
		if temps[i] >= len(w) {
			result++
		}
	}
	return result
}
