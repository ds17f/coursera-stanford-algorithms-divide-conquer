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

func TestUpdateEdges(t *testing.T) {
	testCases := []struct {
		u string
		v string
		t map[string][]string
		e map[string][]string
	}{
		{
			"1",
			"3",
			map[string][]string{
				"1": {"2", "3"},
				"2": {"1", "3", "4"},
				"3": {"1", "2", "4"},
				"4": {"2", "3"},
			},
			map[string][]string{
				"1": {"2", "3"},
				"2": {"3", "3", "4"},
				"3": {"3", "2", "4"},
				"4": {"2", "3"},
			},
		},
	}
	for _, tc := range testCases {
		testCase := tc.t
		expectedResult := tc.e
		algorithms.UpdateEdges(testCase, tc.u, tc.v)
		for k := range testCase {
			sort.Strings(expectedResult[k])
			sort.Strings(testCase[k])
			if expectedResult[k] == nil {
				assert.Nil(t, testCase[k])
			} else {
				assert.Equal(t, strings.Join(expectedResult[k], ","), strings.Join(testCase[k], ","))
			}
		}
	}
}

func TestCollapseNode(t *testing.T) {
	testCases := []struct {
		u string
		v string
		t map[string][]string
		e map[string][]string
	}{
		{
			"1",
			"3",
			map[string][]string{
				"1": {"2", "3"},
				"2": {"3", "3", "4"},
				"3": {"3", "2", "4"},
				"4": {"2", "3"},
			},
			map[string][]string{
				"2": {"3", "3", "4"},
				"3": {"2", "2", "4"},
				"4": {"2", "3"},
			},
		},
		{
			"3",
			"2",
			map[string][]string{
				"2": {"2", "2", "4"},
				"3": {"2", "2", "4"},
				"4": {"2", "2"},
			},
			map[string][]string{
				"2": {"4", "4"},
				"4": {"2", "2"},
			},
		},
	}
	for _, tc := range testCases {
		testCase := tc.t
		expectedResult := tc.e
		algorithms.CollapseNode(testCase, tc.u, tc.v)
		// len should be the same
		assert.Equal(t, len(testCase), len(expectedResult))
		for k := range testCase {
			sort.Strings(expectedResult[k])
			sort.Strings(testCase[k])
			if expectedResult[k] == nil {
				assert.Nil(t, testCase[k])
			} else {
				assert.Equal(t, strings.Join(expectedResult[k], ","), strings.Join(testCase[k], ","))
			}
		}
	}
}
