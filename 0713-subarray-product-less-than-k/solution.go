func numSubarrayProductLessThanK(nums []int, k int) int {
    res := 0

    for i := 0; i < len(nums); i++ {
        p := 1
        for j := i + 0; j < len(nums); j++ {
            p = p * nums[j]
            if p < k {                
                res++
            } else {
                break
            }
        }
    }


    return res
}
