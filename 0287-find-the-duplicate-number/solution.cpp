class Solution {
public:
/*
        This function finds the duplicate number in a list of integers.

        Intuition:
        - The algorithm leverages the fact that the input list contains n+1 integers where each integer is in the range [1, n].
        - By using the indices and values of the list itself, the algorithm effectively creates a cycle that can be detected and utilized to find the duplicate number.
        - Each number is used to mark the index it points to by negating the value at that index. If a number is encountered that points to an already negative index, it means the index is visited twice, indicating the duplicate number.

        Steps:
        - Initialize an index variable `i` to 0.
        - Enter a while loop that continues indefinitely.
        - Inside the loop:
          - If the value at `nums[i]` is negative, it means the index `i` is the duplicate number, so return `i`.
          - Otherwise, negate the value at `nums[i]` and simultaneously update `i` to the value at `nums[i]` (before negation).
*/    
    int findDuplicate(vector<int>& nums) {
        
        int i = 0;
        while (nums[i] > 0) {
            int t = nums[i];
            nums[i] = -1 * nums[i];
            i = t;
        }

        return i;
    }
};
