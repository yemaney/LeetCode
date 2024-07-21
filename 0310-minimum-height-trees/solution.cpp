class Solution
{
public:
    /**
     *  Intuition:
     *  ----------
     *  - The algorithm leverages the properties of trees and the concept of "leaves"
     *    (nodes with only one connection).
     *  - By repeatedly trimming the leaves, the process converges to the centroids of
     *    the graph, which are the roots of MHTs.
     *  - For trees, there are at most two centroids.
     *  
     *  Steps:
     *  ------
     *  1. Handle the edge case where there's only one node (`n == 1`). The root of
     *     the MHT is the single node itself, so return `[0]`.
     *  2. Build an adjacency list `adj` to represent the graph, where `adj[i]` contains
     *     the neighbors of node `i`.
     *  3. Initialize leaves using a deque `leaves` to store all the initial leaves
     *     (nodes with only one connection). Create a dictionary `edgeCnt` to keep track
     *     of the number of connections (edges) for each node.
     *  4. Trim leaves iteratively:
     *     - While there are leaves to process:
     *       - If the number of remaining nodes (`n`) is less than or equal to 2,
     *         return the current leaves as the roots of the MHTs.
     *       - For each leaf, decrement the number of remaining nodes.
     *       - Decrease the edge count of its neighbors.
     *       - If a neighbor becomes a leaf, add it to the `leaves` deque
     */

    vector<int> findMinHeightTrees(int n, vector<vector<int>> &edges)
    {
        if (n == 1)
        {
            return {0};
        }

        unordered_map<int, vector<int>> adj;
        for (auto &&i : edges)
        {
            adj[i[0]].push_back(i[1]);
            adj[i[1]].push_back(i[0]);
        }

        queue<int> que;
        unordered_map<int, int> edgeCnt;
        for (auto &&i : adj)
        {
            if (i.second.size() == 1)
            {
                que.push(i.first);
            }

            edgeCnt[i.first] = i.second.size();
        }

        while (!que.empty())
        {
            if (n <= 2)
            {
                break;
            }

            int size = que.size();

            for (size_t i = 0; i < size; i++)
            {
                int node = que.front();
                que.pop();

                n--;

                for (auto &&i : adj[node])
                {
                    edgeCnt[i]--;
                    if (edgeCnt[i] == 1)
                    {
                        que.push(i);
                    }
                }
            }
        }

        vector<int> ans;
        while (!que.empty())
        {
            ans.push_back(que.front());
            que.pop();
        }

        return ans;
    }
};
