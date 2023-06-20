func moveZeroes(nums []int)  {
    // use l pointer to keep track of index where non-zero element can be put
    // move right through the array
    // if non-zero element is found, place it at index l, increment l
    // loop again to make sure everything after the last non-zero element is zero

    l := 0
    
    for r := range nums {
        if nums[r] != 0 {
            nums[l] = nums[r]
            l++
        }
    }
    for ; l < len(nums); l++ {
        nums[l] = 0
    }
}
