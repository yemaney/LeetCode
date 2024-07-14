class Solution {
public:
    /**
     *   Steps:
     *       1. Create a map (`preMap`) to store course and its prerequisite courses as key-value pairs.
     *       2. For each course, perform a depth-first search (DFS) to traverse the prerequisite chain.
     *       3. Use a set `visited` to track nodes that have been fully processed.
     *       4. Use a set `curCycle` to track nodes currently being visited in the current DFS cycle to detect cycles.
     *       5. If a node is visited twice in the same cycle, a cycle is detected, and the function returns an empty list.
     *       6. If a course with no prerequisites is found, it is appended to the result list `ans`.
     *       7. If a course with every prerequisite already processed is found, it is appended to the result list `ans`.
     *   
     *  Example:
     *       prerequisites = [[1,0], [0,2], [1,2]]
     *       
     *       The function should return a list such as [2, 0, 1] or [2, 1, 0].
     *
     *       Explanation:
     *           preMap = {
     *               0: [2],
     *               1: [0, 2],
     *               2: []
     *           }
     */
    vector<int> findOrder(int numCourses, vector<vector<int>>& prerequisites) {

        unordered_map<int,vector<int>> preMap;
        for (auto &&c : prerequisites)
        {
            preMap[c[0]].push_back(c[1]);
        }

        vector<int> ans;
        unordered_set<int> processed;
        unordered_set<int> cycle;

        for (size_t i = 0; i < numCourses; i++)
        {
            if (!dfs(i, preMap, processed, cycle, ans))
            {
                return {};
            }
            
        }
        
        return ans;
    }
private:
    bool dfs(int c, unordered_map<int, vector<int>>& preMap, unordered_set<int>& processed, unordered_set<int>& cycle, vector<int>& ans) {
        if (cycle.find(c) != cycle.end())
        {
            return false;
        }
        if (processed.find(c) != processed.end())
        {
            return true;
        }
        
        cycle.insert(c);
        // The first time this passes as true, is when `c` doesn't have any preqs.
        // All passes occur if c has no preqs or if they have all been already been processed.
        for (auto &&p : preMap[c])
        {
            if (!dfs(p, preMap, processed, cycle, ans))
            {
                return false;
            }
            
        }
        
        cycle.erase(c);
        processed.insert(c);
        ans.push_back(c);

        return true;
    }
};
