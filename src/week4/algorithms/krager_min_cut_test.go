package algorithms_test

import (
	"strings"
	"testing"

	"github.com/damiansilbergleithcunniff/algorithms-divide-conquer/src/week4/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestKragerMinCut(t *testing.T) {
}

func TestReplaceReference(t *testing.T) {
	testCase := [][]string{
		{"2", "5"},
		{"1", "3", "4"},
		{"2", "4"},
		{"2", "3", "5"},
		{"1", "4"},
	}
	expectedResult := [][]string{
		{"1", "5"},
		{"1", "3", "4"},
		{"1", "4"},
		{"1", "3", "5"},
		{"1", "4"},
	}

	algorithms.ReplaceReference(testCase, "2", "1")

	for i := range testCase {
		assert.Equal(t, strings.Join(expectedResult[i], ","), strings.Join(testCase[i], ","))
	}
}
