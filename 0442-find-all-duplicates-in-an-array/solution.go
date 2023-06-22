func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func findDuplicates(nums []int) []int {
    // loop through array, for each value v use (v-1) as an index
    // turn the element in the array at the index into a negative number
    // if its already negative then its a duplice then add it to res

    res := []int{}

    for _, v := range nums {
        p :=  abs(v)

        if nums[p-1] < 0 {
            res = append(res, int(math.Abs(float64(v))))
            continue
        } 
        nums[p-1] *= -1
    }

    return res
}
