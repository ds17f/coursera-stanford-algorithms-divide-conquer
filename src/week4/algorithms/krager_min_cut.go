package algorithms

import "strconv"

// KragerMinCut implements Krager's randomized algorithm
// for finding the minimum number of cuts in a graph.
// input should be an adjacency list and the element at
// position n should represent the node n+1.  So the 10th
// entry in the array is the 11th numbered vertex
func KragerMinCut(input [][]string) int {
	return 0
}

// ReplaceReference replaces all the references to the node `from`
// with the node `to` in the input adjacency list
func ReplaceReference(input [][]string, from string, to string) {
	fromI, _ := strconv.Atoi(from)
	// This is the node that we are collapsing
	// so we need to iterate through it and then
	// replace all references to it with the node we're
	// collapsing into
	nodeToChange := input[fromI-1]
	for _, v := range nodeToChange {
		// each edge that the node touches needs to be examined
		vI, err := strconv.Atoi(v)
		if err == nil {
			adjacentNode := input[vI-1]
			// now iterate through the adjacentNode and replace
			// the collapsed node with the node we collapse into.
			for i, w := range adjacentNode {
				if w == from {
					adjacentNode[i] = to
				}
			}

		}
	}
}
