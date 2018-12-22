package algorithms

// KargerMinCut implements Karger's randomized algorithm
// for finding the minimum number of cuts in a graph.
// input is a map adjacency list with the key as the node id
// and the value the list of verticies that make up that node's edges
func KargerMinCut(input map[string][]string) int {
	return 0
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
