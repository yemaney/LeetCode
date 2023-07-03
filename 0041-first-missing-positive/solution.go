func firstMissingPositive(nums []int) int {
    fmt.Println(nums)
    // turn negatives into 0, to remove false flags
    for i, v := range nums {
        if v < 0 {
            nums[i] = 0
        }
    }
    fmt.Println(nums)
    // flag with negatives
    for _, v := range nums {
        av := abs(v)
        if av == 0 || av > len(nums) {
            continue
        }
        nums[av-1] = flag(nums[av-1], len(nums))
    }
    fmt.Println(nums)
    ans := 1
    for _, v := range nums {
        if v >= 0 {
            break
        }
        ans++
    }
    return ans
}

func abs(n int) int {
    if n < 0 {
        return -1 * n
    }
    return n
}

func flag(n int, lim int) int {
    if n > 0 {
        return -1 * n
    } else if n == 0 {
        return -1 * (lim + 1)
    } else {
    return n
    }
}
