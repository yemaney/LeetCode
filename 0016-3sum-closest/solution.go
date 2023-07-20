func threeSumClosest(nums []int, target int) int {
    res := nums[0] + nums[1] + nums[2]

    for i := 0; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums) - 1; j++ {
            for k := j+1; k < len(nums); k++ {
                curRes := nums[i] + nums[j] + nums[k]
                if abs(target - curRes) < abs(target - res) {
                    res = curRes
                }
            }
        }
    
    }
    return res
}

func abs(i int) int {
    if i < 0 {
        return -i
    }
    return i
}
