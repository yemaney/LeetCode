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
class Solution {
public:
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
    vector<int> rightSideView(TreeNode* root) {
        if (root == NULL)
        {
            return {};
        }

        queue<TreeNode*> que;
        que.push(root);
        vector<int> ans;

        while (!que.empty())
        {
            int level_right;
            int level_size = que.size();

            for (size_t i = 0; i < level_size; i++)
            {
                TreeNode* node = que.front();
                que.pop();
                level_right = node->val;

                if (node->left != NULL)
                {
                    que.push(node->left);
                }
                if (node->right != NULL)
                {
                    que.push(node->right);
                }
            }

            ans.push_back(level_right);
        }

        return ans;
    }
};
