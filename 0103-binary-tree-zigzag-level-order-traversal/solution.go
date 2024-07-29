/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/*
 *	Perform a level order traversal on a binary tree and reverse the order of nodes at every alternate level.
 *          	 		
 *	Intuition:
 *	----------
 *	The goal is to traverse a binary tree in level order, but with a twist: we want to reverse the order of the levels. 
 *	To achieve this, we use a queue to facilitate level-order traversal and a flag (`rev`) to determine whether to reverse 
 *	the current level's nodes before adding them to the result list.
 *   	        	
 *	Steps:
 *	------
 *	1. Initialization:
 *	    - Initialize a flag `rev` to `False`. This flag will help decide whether to reverse the current level's nodes.
 *	    - Initialize an empty list `ans` to store the final result.
 *	    - If the root is `None`, return `ans` immediately as the tree is empty.
 *	    - Initialize a deque `que` and append the root node to it.
 *	2. Level Order Traversal:
 *	    - While the queue is not empty:
 *	        1. Initialize an empty list `level` to store nodes of the current level.
 *	        2. Iterate over the current size of the queue:
 *	            - Pop a node from the left side of the queue.
 *	            - Append its value to the `level` list.
 *	            - If the node has a left child, append it to the queue.
 *	            - If the node has a right child, append it to the queue.
 *	        3. If the `rev` flag is `True`, reverse the `level` list.
 *	        4. Append the `level` list to `ans`.
 *	        5. Toggle the `rev` flag for the next level.
 *	3. Return the Result:
 *	    - Return the `ans` list containing the nodes in the required order.
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}

	rev := false
	que := []*TreeNode{}
	que = append(que, root)

	for len(que) > 0 {
		level := []int{}
		level_size := len(que)

		for i := 0; i < level_size; i++ {
			node := que[0]
			que = que[1:]
			level = append(level, node.Val)


			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}

		if rev {
			slices.Reverse(level)
		}
		ans = append(ans, level)
		rev = !rev
	}

	return ans
}
