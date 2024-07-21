class Solution:
    def findMinHeightTrees(self, n : int, edges : list[list[int]]) -> list[int]:
        """
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
        """



        if n == 1:
            return [0]

        adj : dict[int, list[int]] = {i:[] for i in range(n)}
        for n1, n2 in edges:
            adj[n1].append(n2)
            adj[n2].append(n1)
        


        edgeCnt = {}
        leaves : deque[int] = deque()
        for src, nb in adj.items():
            if len(nb) == 1:
                leaves.append(src)
            edgeCnt[src] = len(nb)

        while leaves:
            if n <= 2:
                return list(leaves)
            for _ in range(len(leaves)):
                node = leaves.popleft()
                n -= 1
                for nb in adj[node]:
                    edgeCnt[nb] -= 1
                    if edgeCnt[nb] == 1:
                        leaves.append(nb)
