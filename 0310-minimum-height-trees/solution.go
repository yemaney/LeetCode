func findMinHeightTrees(n int, edges [][]int) []int {
	/*
	   Intuition:
	   ----------
	   - The algorithm leverages the properties of trees and the concept of "leaves"
	     (nodes with only one connection).
	   - By repeatedly trimming the leaves, the process converges to the centroids of
	     the graph, which are the roots of MHTs.
	   - For trees, there are at most two centroids.

	   Steps:
	   ------
	   1. Handle the edge case where there's only one node (`n == 1`). The root of
	      the MHT is the single node itself, so return `[0]`.
	   2. Build an adjacency list `adj` to represent the graph, where `adj[i]` contains
	      the neighbors of node `i`.
	   3. Initialize leaves using a deque `leaves` to store all the initial leaves
	      (nodes with only one connection). Create a dictionary `edgeCnt` to keep track
	      of the number of connections (edges) for each node.
	   4. Trim leaves iteratively:
	      - While there are leaves to process:
	        - If the number of remaining nodes (`n`) is less than or equal to 2,
	          return the current leaves as the roots of the MHTs.
	        - For each leaf, decrement the number of remaining nodes.
	        - Decrease the edge count of its neighbors.
	        - If a neighbor becomes a leaf, add it to the `leaves` deque
	*/
	if n == 1 {
		return []int{0}
	}

	adj := map[int][]int{}
	for _, v := range edges {
		adj[v[0]] = append(adj[v[0]], v[1])
		adj[v[1]] = append(adj[v[1]], v[0])
	}

	que := []int{}
	edgeCnt := map[int]int{}
	for k, v := range adj {
		if len(v) == 1 {
			que = append(que, k)
		}
		edgeCnt[k] = len(v)
	}

	for len(que) > 0 {
		if n <= 2 {
			break
		}

		size := len(que)
		for i := 0; i < size; i++ {
			node := que[0]
			que = que[1:]
			n--

			for _, v := range adj[node] {
				edgeCnt[v]--
				if edgeCnt[v] == 1 {
					que = append(que, v)
				}
			}
		}

	}

	return que
}

