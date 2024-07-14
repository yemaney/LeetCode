class Solution {
public:
    /**
     *   1. create map of course : preqs pairs
     *   2. for each course do dfs to traverse the preq chain
     *   3. store visited nodes in set along the way
     *   4. if a node is visited twice exit early as False, else True
     *
     *   This can be done because of the property `0 <= ai, bi < numCourses`. 
     *   The values for courses and preqs being within the bounds of numCourses
     *   enables hopping through preMap as a tree.
     *
     *   Example:
     *       prerequisites = [[1,0], [0,2], [1,2]]
     *
     *       preMap = {
     *           0 : [2],
     *           1 : [0, 2],
     *           2 : []
     *       }
     */
    bool canFinish(int numCourses, vector<vector<int>>& prerequisites) {

        unordered_map<int, vector<int>> preMap;
        for (size_t i = 0; i < prerequisites.size(); i++)
        {
            preMap[prerequisites[i][0]].push_back(prerequisites[i][1]);
        }

        unordered_set<int> visited;

        for (size_t i = 0; i < numCourses; i++)
        {
            if (dfs(i, preMap, visited) == false)
            {
                return false;
            }            
        }

        return true; 
    }
private:
    bool dfs(int x, unordered_map<int, vector<int>>& preMap, unordered_set<int>& visited) {
        if (visited.find(x) != visited.end())
        {
            return false;
        }
        
        if (preMap[x].size() == 0)
        {
            return true;
        }

        visited.insert(x);
        for (size_t i = 0; i < preMap[x].size(); i++)
        {
            if (!dfs(preMap[x][i], preMap, visited))
            {
                return false;
            }
        }

        visited.erase(x);
        preMap[x] = vector<int>{};
        
        return true;
    }
};
