/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


/**
	Given the root of a binary tree, return the values of the nodes
	you can see ordered from top to bottom when looking from the right side.

	Intuition:
	----------
	- Use a breadth-first search (BFS) to traverse the tree level by level.
	- For each level, the last node encountered is the rightmost node visible
		from that level when viewed from the right side.

	Steps:
	-----
	1. Initialize a que for BFS and append the root node to it.
	2. Initialize an empty list to store the right side view values.
	3. While the que is not empty, repeat the following:
	- Determine the number of nodes at the current level (level_size).
	- Iterate through all nodes at the current level:
	- Pop the leftmost node from the que.
	- Update the rightmost node's value for the current level.
	- Append the left and right children of the node to the que if they exist.
	- Append the rightmost node's value to the right side view list.

	4. Return the list containing the right side view values.
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	que := []*TreeNode{}
	que = append(que, root)
	ans := []int{}

	for len(que) > 0 {
		var level_right int
		level_size := len(que)

		for i := 0; i < level_size; i++ {
			node := que[0]
			que = que[1:]
			level_right = node.Val

			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}

		ans = append(ans, level_right)
	}

	return ans
}

