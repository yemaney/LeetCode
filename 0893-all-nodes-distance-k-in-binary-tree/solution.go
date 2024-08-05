/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
    Finds all nodes at distance K from the target node in a binary tree.

    Intuition:
    - Convert the binary tree into an adjacency list to handle it like an undirected graph.
    - Use Breadth-First Search (BFS) starting from the target node to find all nodes at distance K.

    Steps:
    1. Check if the root is None. If so, return an empty list.
    2. Initialize the adjacency list using the binary tree.
    3. Initialize a queue with the target node's value and a visited set containing the target node's value.
    4. Perform BFS until the distance equals K:
        - For each node at the current level, add its neighbors to the queue if they haven't been visited.
        - Increment the steps counter.
    5. Return the list of nodes in the queue.
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
    if root == nil {
        return []int{}
    }

    adj := map[int][]int{}
    init_adj(root, adj)
    
    que := []int{target.Val}
    visited := map[int]struct{}{target.Val: struct{}{}}
    steps := 0

    for len(que) > 0 {
        if steps == k {
            break
        }

        n_size := len(que)
        for i := 0; i < n_size; i++{
            node := que[0]
            que = que[1:]

            for _, v := range adj[node] {
                if _, ok := visited[v]; !ok {
                    que = append(que, v)
                    visited[v] = struct{}{}
                }
            }
        }
        
        steps += 1
    }


    return que
}


/*
    Initializes the adjacency list for the binary tree.

    Intuition:
    - Traverse the binary tree and create bidirectional edges in the adjacency list for each node.

    Steps:
    1. If the node has a left child, add bidirectional edges between the node and the left child.
    2. Recursively initialize the adjacency list for the left child.
    3. If the node has a right child, add bidirectional edges between the node and the right child.
    4. Recursively initialize the adjacency list for the right child.
*/
func init_adj(node *TreeNode, adj map[int][]int) {
    if _, ok := adj[node.Val]; !ok {
        adj[node.Val] = []int{} 
    }

    if node.Left != nil {
        adj[node.Val] = append(adj[node.Val], node.Left.Val)
        adj[node.Left.Val] = []int{node.Val}
        
        init_adj(node.Left, adj)
    }

    if node.Right != nil {
        adj[node.Val] = append(adj[node.Val], node.Right.Val)
        adj[node.Right.Val] = []int{node.Val}
        
        init_adj(node.Right, adj)
    }
}
