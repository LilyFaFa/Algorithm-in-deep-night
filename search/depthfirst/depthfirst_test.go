package depthfirst

import (
	"testing"

	"github.com/LilyFaFa/Algorithm-in-deep-night/collections"
)

func TestAccessNode(t *testing.T) {
	// 10 * 10 邻接矩阵
	testCase := [][]int{
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0}, //0
		{1, 0, 1, 0, 0, 1, 0, 0, 0, 0}, //1
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 0}, //2
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 1}, //3
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, //4
		{0, 1, 0, 0, 0, 0, 1, 0, 1, 0}, //5
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0}, //6
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0}, //7
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 1}, //8
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0}, //9
	}
	currentNode := 0
	actual := &collections.Stack{}

	expected := &collections.Stack{}
	//http://golang.org/doc/faq#convert_slice_of_interface
	//将[]int转化为[]intereface{}是不可以直接转换的
	expectedList := []int{1, 3}
	s := make([]interface{}, len(expectedList))
	for i, v := range expectedList {
		s[i] = v
	}
	expected.Items = s

	visited := []bool{}
	for i := 0; i < len(testCase); i++ {
		visited = append(visited, false)
	}
	visited[0] = true
	AccessNode(testCase, currentNode, actual, visited)

	if len(actual.Items) != len(expected.Items) {
		t.Errorf("expected lenth=%v, actual lenth=%v", len(expected.Items), len(actual.Items))
	} else {
		for i, _ := range actual.Items {
			itemAc := actual.Items[i].(int)
			itemEx := expected.Items[i].(int)
			if itemAc != itemEx {
				t.Errorf("expected=%v, actual=%v", itemEx, itemAc)
			}
		}
	}
}

func TestDepthFirst(t *testing.T) {
	// 10 * 10 邻接矩阵
	testCase := [][]int{
		{0, 1, 0, 1, 0, 0, 0, 0, 0, 0}, //0
		{1, 0, 1, 0, 0, 1, 0, 0, 0, 0}, //1
		{0, 1, 0, 1, 1, 0, 0, 0, 0, 0}, //2
		{1, 0, 1, 0, 0, 0, 0, 0, 0, 1}, //3
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, //4
		{0, 1, 0, 0, 0, 0, 1, 0, 1, 0}, //5
		{0, 0, 0, 0, 0, 1, 0, 1, 0, 0}, //6
		{0, 0, 0, 0, 1, 0, 1, 0, 0, 0}, //7
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 1}, //8
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0}, //9
	}

	expected := collections.Queue{}
	//http://golang.org/doc/faq#convert_slice_of_interface
	//将[]int转化为[]intereface{}是不可以直接转换的
	expectedItems := []int{0, 3, 9, 8, 5, 6, 7, 4, 2, 1}
	s := make([]interface{}, len(expectedItems))
	for i, v := range expectedItems {
		s[i] = v
	}
	expected.Items = s

	actual := DepthFirst(testCase)
	if len(actual.Items) != len(expected.Items) {
		t.Errorf("expected lenth=%v, actual lenth=%v", len(expected.Items), len(actual.Items))
	} else {
		for i, _ := range expected.Items {
			itemAc := actual.Items[i].(int)
			itemEx := expected.Items[i].(int)
			if itemAc != itemEx {
				t.Errorf("expected=%v, actual=%v", itemEx, itemAc)
			}
		}
	}

}
