/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	/*
    1. traverse tree by level
    2. increment the ans each level
    3. exit as soon as a root leaf is found
	*/

	ans := 0

	if root == nil {
		return ans
	}

	q := []*TreeNode{}
	q = append(q, root)

	for len(q) > 0 {
		ans += 1
		size := len(q)

		for i := 0; i < size; i++ {
            node := q[0]
            q = q[1:]

            if node.Left == nil && node.Right == nil {
				return ans
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

	}

	return ans
}
