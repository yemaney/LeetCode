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
/*
 *	Perform a level order traversal on a binary tree and reverse the order of nodes at every alternate level.
 *          	 		
 *	Intuition:
 *	----------
 *	The goal is to traverse a binary tree in level order, but with a twist: we want to reverse the order of the levels. 
 *	To achieve this, we use a queue to facilitate level-order traversal and a flag (`rev`) to determine whether to reverse 
 *	the current level's nodes before adding them to the result vector.
 *   	        	
 *	Steps:
 *	------
 *	1. Initialization:
 *	    - Initialize a flag `rev` to `false`. This flag will help decide whether to reverse the current level's nodes.
 *	    - Initialize an empty vector `ans` to store the final result.
 *	    - If the root is `None`, return `ans` immediately as the tree is empty.
 *	    - Initialize a `que` and append the root node to it.
 *	2. Level Order Traversal:
 *	    - While the queue is not empty:
 *	        1. Initialize an empty vector `level` to store nodes of the current level.
 *	        2. Iterate over the current size of the queue:
 *	            - Pop a node from the left side of the queue.
 *	            - Append its value to the `level` vector.
 *	            - If the node has a left child, append it to the queue.
 *	            - If the node has a right child, append it to the queue.
 *	        3. If the `rev` flag is `True`, reverse the `level` vector.
 *	        4. Append the `level` vector to `ans`.
 *	        5. Toggle the `rev` flag for the next level.
 *	3. Return the Result:
 *	    - Return the `ans` vector containing the nodes in the required order.
*/
    vector<vector<int>> zigzagLevelOrder(TreeNode* root) {
        vector<vector<int>> ans;
        if (!root)
        {
            return ans;
        }

        bool rev = false;
        queue<TreeNode*> que;
        que.push(root);

        while (!que.empty())
        {
            vector<int> level;
            int level_size = que.size();

            for (size_t i = 0; i < level_size; i++)
            {
                TreeNode* node = que.front();
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

            if (rev)
            {
                reverse(level.begin(), level.end());
            }
            ans.push_back(level);
            rev = !rev;
            
        }
        
        return ans;
    }
};
