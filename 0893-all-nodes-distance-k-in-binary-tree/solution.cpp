/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
class Solution
{
public:
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
    vector<int> distanceK(TreeNode *root, TreeNode *target, int k)
    {
        if (root == NULL)
        {
            return {};
        }

        unordered_map<int, vector<int>> adj;
        init_adj(root, adj);

        queue<int> que;
        que.push(target->val);
        unordered_set<int> visited;
        visited.insert(target->val);
        int steps = 0;

        while (!que.empty())
        {
            if (steps == k)
            {
                break;
            }

            int q_size = que.size();
            for (size_t i = 0; i < q_size; i++)
            {
                int node = que.front();
                que.pop();

                for (auto &&n : adj[node])
                {
                    if (visited.find(n) == visited.end())
                    {
                        que.push(n);
                        visited.insert(n);
                    }
                }
            }
            steps++;
        }

        vector<int> ans;
        while (!que.empty())
        {
            ans.push_back(que.front());
            que.pop();
        }

        return ans;
    }

private:
/**
    Initializes the adjacency list for the binary tree.

    Intuition:
    - Traverse the binary tree and create bidirectional edges in the adjacency list for each node.

    Steps:
    1. If the node has a left child, add bidirectional edges between the node and the left child.
    2. Recursively initialize the adjacency list for the left child.
    3. If the node has a right child, add bidirectional edges between the node and the right child.
    4. Recursively initialize the adjacency list for the right child.

 */
    void init_adj(TreeNode *node, unordered_map<int, vector<int>> &adj)
    {
        if (adj.find(node->val) == adj.end())
        {
            adj[node->val] = {};
        }

        if (node->left != NULL)
        {
            adj[node->val].push_back(node->left->val);
            adj[node->left->val].push_back(node->val);

            init_adj(node->left, adj);
        }

        if (node->right != NULL)
        {
            adj[node->val].push_back(node->right->val);
            adj[node->right->val].push_back(node->val);

            init_adj(node->right, adj);
        }
    }
};
