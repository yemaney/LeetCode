func sortedSquares(nums []int) []int {
    arr := make([]int, len(nums))

    l, r := 0, len(nums) - 1
    p := len(nums)-1 
    for l <= r {
        if abs(nums[l]) > abs(nums[r]) {
            arr[p] = nums[l] * nums[l]
            l++
        } else {
            arr[p] = nums[r] * nums[r] 
            r--
        }
        p--
    }

    return arr
}

func abs(n int) int {
    if n < 0 {
        return -1 * n
    }
    return n
}
