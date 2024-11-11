class Solution {
public:
	/*
		Notes
		-----
		The function follows these steps:
		- Verify if the product `m * n` matches the length of `original`.
		- Return an empty list if the total number of elements does not match.
		- Initialize an empty list `ans` to store the final 2D array.
		- Use a loop to iterate through `original` and build subarrays of length `n`.
		- Append each completed subarray to `ans` and reset `subarray` for the next row.
		- Return `ans` after processing all elements.
	*/
    vector<vector<int>> construct2DArray(vector<int>& original, int m, int n) {
        vector<vector<int>> ans;
        if (original.size() != m * n)
        {
            return ans;
        }

        int i = 0;
        vector<int> subarray;
        while (i < original.size())
        {
            subarray.push_back(original[i]);
            i++;

            if (subarray.size() == n)
            {
                ans.push_back(subarray);
                subarray = {};
            }
        }
        
        return ans;
    }
};
