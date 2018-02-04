package lcs

import (
	"testing"
)

func TestLCS(t *testing.T) {
	expectedSeq := "BSGSG"
	expectedLen := 5
	testCaseX := "BFRSDFGSFG"
	testCaseY := "ABSGRGSRGXS"
	actualSeq, actualLen := LCS(testCaseX, testCaseY)
	if expectedLen != actualLen {
		t.Errorf("lenth is error,expected=%v,actual=%v", expectedLen, actualLen)
	}
	if expectedSeq != actualSeq {
		t.Errorf("subsequence is error,expected=%v,actual=%v", expectedSeq, actualSeq)
	}
}
