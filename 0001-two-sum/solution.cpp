class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        // 1. iterate over the vector 
        // 2. check if the required number to make target with current number exists
        //    in map and return the indices as result
        // 3. otherwise store visited number in map as key and the index as value 
    
        vector<int> result = {};
        unordered_map<int, int> m;

        for (int i=0; i < nums.size(); ++i) {
            if (m.find(target-nums[i]) != m.end()) {
                result = {i, m[target-nums[i]]};
                return result;
            }
            m[nums[i]] = i;
        }
        return result;
    }
};
