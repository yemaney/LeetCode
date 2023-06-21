func findDuplicate(nums []int) int {
    
    i := 0
    for ; true ; {
        if nums[i] < 0 {
            return i
        }
        nums[i], i = -1 * nums[i], nums[i] 
        
    }
    return 0
}
