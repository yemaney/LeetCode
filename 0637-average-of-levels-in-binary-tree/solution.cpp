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
    vector<double> averageOfLevels(TreeNode* root) {
        vector<double> ans;
        if (root == nullptr) {
            return ans;
        }

        queue<TreeNode*> q;
        q.push(root);

        while (!q.empty())
        {
            int64_t  sum = 0;
            int size = q.size();

            for (size_t i = 0; i < size; i++)
            {
                TreeNode* node = q.front();
                q.pop();

                sum = sum + node->val;

                if (node->left != nullptr) {
                    q.push(node->left);
                }
                if (node->right != nullptr) {
                    q.push(node->right);
                }
            }

            double avg = static_cast<double>(sum) / static_cast<double>(size);
            ans.push_back(avg);
        }


        return ans; 

    }
};
