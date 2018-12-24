package algorithms

import (
	"math"
	"math/rand"
	"time"
)

// KargerMinCut implements Karger's randomized algorithm
// for finding the minimum number of cuts in a graph.
// input is a map adjacency list with the key as the node id
// and the value the list of verticies that make up that node's edges
func KargerMinCut(input map[string][]string) int {
	n := float64(len(input))
	iterations := int(math.Ceil((n * n) * math.Ln10 * n))
	iterations = int(math.Min(n*n, 1000.0))
	min := 999999
	for i := 0; i < iterations; i++ {
		clonedInput := cloneMap(input)
		result := RunKargerMinCut(clonedInput)
		if result < min {
			min = result
		}
	}
	return min
}

// cloneSlice creates a clone of the slice
// in a new array and then returns a slice
// of that new array
func cloneSlice(slice []string) []string {
	clone := make([]string, len(slice))
	copy(clone, slice)
	return clone
}

func cloneMap(input map[string][]string) map[string][]string {
	clone := make(map[string][]string)
	for k, v := range input {
		clone[k] = cloneSlice(v)
	}
	return clone
}

// RunKargerMinCut does a single iteration of the Karger
// algorithm, collapsing the graph until there are 2 nodes
func RunKargerMinCut(input map[string][]string) int {
	for len(input) > 2 {
		u, v := ChooseRandomEdge(input)
		UpdateEdges(input, u, v)
		CollapseNode(input, u, v)
	}
	for _, val := range input {
		return len(val)
	}
	return -1
}

// ChooseRandomEdge will pick a random element in the map
// then a random element from the list held in the first choice
func ChooseRandomEdge(input map[string][]string) (string, string) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	// get an slice of the keys so that we can pick one at random
	keys := make([]string, 0, len(input))
	for k := range input {
		keys = append(keys, k)
	}
	u := keys[r1.Intn(len(input))]
	uEdges := input[u]
	v := uEdges[r1.Intn(len(uEdges))]
	return u, v
}

// UpdateEdges replaces all references to "u" with references to "v"
func UpdateEdges(input map[string][]string, u string, v string) {
	collapseFrom := input[u]
	// iterate through all the edges in the node we are collapsing
	for _, adjacentVertex := range collapseFrom {
		adjacentNode := input[adjacentVertex]
		// now iterate through all the edges of the adjacent node
		// so we can locate the collapsing vertex (u) and
		// update them to point to the node we're collapsing into (v)
		for i, w := range adjacentNode {
			if w == u {
				adjacentNode[i] = v
			}
		}
	}
}

// CollapseNode moves all verticies from u -> v, then removes
// the node at u from the input
func CollapseNode(input map[string][]string, u string, v string) {
	fromNode := input[u]
	toNode := input[v]

	var newNodeList []string
	for _, vertex := range toNode {
		if vertex != v {
			newNodeList = append(newNodeList, vertex)
		}
	}
	for _, vertex := range fromNode {
		if vertex != v {
			newNodeList = append(newNodeList, vertex)
		}
	}

	input[v] = newNodeList
	delete(input, u)
}
