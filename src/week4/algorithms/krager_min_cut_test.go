package algorithms_test

import (
	"sort"
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
		sort.Strings(expectedResult[i])
		sort.Strings(testCase[i])
		assert.Equal(t, strings.Join(expectedResult[i], ","), strings.Join(testCase[i], ","))
	}
}

func TestReplaceReference2(t *testing.T) {
	testCase := [][]string{
		{"2", "3"},
		{"1", "3", "4"},
		{"1", "2", "4"},
		{"2", "3"},
	}
	expectedResult := [][]string{
		{"2", "3"},
		{"3", "3", "4"},
		{"3", "2", "4"},
		{"2", "3"},
	}

	algorithms.ReplaceReference(testCase, "1", "3")

	for i := range testCase {
		sort.Strings(expectedResult[i])
		sort.Strings(testCase[i])
		assert.Equal(t, strings.Join(expectedResult[i], ","), strings.Join(testCase[i], ","))
	}
}
func TestCopyToTarget(t *testing.T) {
	testCase := [][]string{
		{"1", "5"},
		{"1", "3", "4"},
		{"1", "4"},
		{"1", "3", "5"},
		{"1", "4"},
	}
	expectedResult := [][]string{
		{"3", "4", "5"},
		nil,
		{"1", "4"},
		{"1", "3", "5"},
		{"1", "4"},
	}

	algorithms.CopyToTarget(testCase, "2", "1")

	for i := range testCase {
		sort.Strings(expectedResult[i])
		sort.Strings(testCase[i])
		if expectedResult[i] == nil {
			assert.Nil(t, testCase[i])
		} else {
			assert.Equal(t, strings.Join(expectedResult[i], ","), strings.Join(testCase[i], ","))
		}
	}
}
func TestCopyToTarget2(t *testing.T) {
	testCases := []struct {
		u string
		v string
		t [][]string
		e [][]string
	}{
		{
			"1",
			"3",
			[][]string{
				{"2", "3"},
				{"3", "3", "4"},
				{"3", "2", "4"},
				{"2", "3"},
			},
			[][]string{
				nil,
				{"3", "3", "4"},
				{"2", "2", "4"},
				{"2", "3"},
			},
		},
		{
			"3",
			"2",
			[][]string{
				nil,
				{"2", "2", "4"},
				{"2", "2", "4"},
				{"2", "2"},
			},
			[][]string{
				nil,
				{"4", "4"},
				nil,
				{"2", "2"},
			},
		},
	}

	for _, tc := range testCases {
		testCase := tc.t
		expectedResult := tc.e
		algorithms.CopyToTarget(testCase, tc.u, tc.v)
		for i := range testCase {
			sort.Strings(expectedResult[i])
			sort.Strings(testCase[i])
			if expectedResult[i] == nil {
				assert.Nil(t, testCase[i])
			} else {
				assert.Equal(t, strings.Join(expectedResult[i], ","), strings.Join(testCase[i], ","))
			}
		}
	}
}
