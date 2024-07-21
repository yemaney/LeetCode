/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	/*
	   Intuition:
	   ----------
	   - Use a queue to perform a level order traversal.
	   - Insert each level's nodes at the beginning of the result list to achieve bottom-up order.

	   Steps:
	   ------
	   1. Initialize an empty result list `res`.
	   2. If the root is None, return the empty result list.
	   3. Use a queue to perform level order traversal starting with the root.
	   4. For each level:
	       - Initialize an empty list `level` to store nodes of the current level.
	       - Process all nodes at the current level.
	       - Append the value of each node to the `level` list.
	       - Add the left and right children of each node to the queue for the next level.
	       - Insert the `level` list to the result list `res`.
	   5. Return the reversed result list `res`.
	*/
	res := [][]int{}

	if root == nil {
		return res
	}

	que := []*TreeNode{}
	que = append(que, root)

	for len(que) > 0 {
		level := []int{}
		size := len(que)

		for i := 0; i < size; i++ {
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
		res = append(res, level)
	}

	slices.Reverse(res)

	return res
}
