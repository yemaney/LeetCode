class Solution {
public:
    int peakIndexInMountainArray(vector<int>& arr) {
        int left = 0;
        int right = arr.size() - 1;
        while ((left <= right) && (left >= 0) && (right < arr.size())){
            int mid = left + (right - left) / 2;

            std::cout << left << right << mid << std::endl;;

            if ((arr[mid-1] < arr[mid]) && (arr[mid+1] < arr[mid])) {
                return mid; // Found the target at index mid
            } else if (arr[mid] < arr[mid+1]) {
                left = mid; // Search the right half
            } else {
                right = mid; // Search the left half
            }
        }
        return -1;
    }
};



