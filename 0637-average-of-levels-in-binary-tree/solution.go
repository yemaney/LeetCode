/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	ans := []float64{}
	if root == nil {
		return ans
	}

	q := []*TreeNode{}
    q = append(q, root)

	for len(q) > 0 {
		sum := 0
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			sum += node.Val
			
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		avg := float64(sum) / float64(size)
		ans = append(ans, avg)

	}

	return ans
}
