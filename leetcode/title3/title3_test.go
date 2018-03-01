package title3

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testcase := []string{
		"abcabcbbb",
		"abcacd",
		"a",
		"",
	}
	expected := []int{
		3,
		3,
		1,
		0,
	}
	actual := []int{}
	for i := 0; i < len(testcase); i++ {
		actual = append(actual, lengthOfLongestSubstring(testcase[i]))
	}
	for i, a := range actual {
		if a != expected[i] {
			t.Errorf("expected=%v, actual=%v", expected[i], a)
		}
	}

}
