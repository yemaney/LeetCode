/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution
{
public:
    /**
	 *  Intuition:
	 *  ----------
	 *  - Use a queue to perform a level order traversal.
	 *  - Insert each level's nodes at the beginning of the result list to achieve bottom-up order.
     * 
	 *  Steps:
	 *  ------
	 *  1. Initialize an empty result list `res`.
	 *  2. If the root is None, return the empty result list.
	 *  3. Use a queue to perform level order traversal starting with the root.
	 *  4. For each level:
	 *      - Initialize an empty list `level` to store nodes of the current level.
	 *      - Process all nodes at the current level.
	 *      - Append the value of each node to the `level` list.
	 *      - Add the left and right children of each node to the queue for the next level.
	 *      - Insert the `level` list to the result list `res`.
	 *  5. Return the reversed result list `res`.
     */
    vector<vector<int>> levelOrderBottom(TreeNode *root)
    {
        vector<vector<int>> res;

        if (!root)
        {
            return res;
        }

        queue<TreeNode *> que;
        que.push(root);

        while (!que.empty())
        {
            vector<int> level;
            int size = que.size();

            for (size_t i = 0; i < size; i++)
            {
                TreeNode *node = que.front();
                que.pop();
                level.push_back(node->val);

                if (node->left)
                {
                    que.push(node->left);
                }
                if (node->right)
                {
                    que.push(node->right);
                }
            }
            res.push_back(level);
        }

        reverse(res.begin(), res.end());
        return res;
    }
};
