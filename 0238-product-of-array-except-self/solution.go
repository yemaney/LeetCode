func productExceptSelf(nums []int) []int {
    // l    [1     a    ab    abc]
    // r    [bcd  cd     d     1 ]
    // res  [bcd   acd   abd   abc]

    l := make([]int, len(nums))
    r := make([]int, len(nums))
    res := make([]int, len(nums))

    l[0] = 1
    for i := 1; i < len(nums); i++ {
        l[i] = nums[i-1] * l[i-1]
    }

    r[len(nums)-1] = 1
    for i := len(nums)-2; i >= 0; i-- {
        r[i] = nums[i+1] * r[i+1]
    }

    for i := 0; i < len(nums); i++{
        res[i] = l[i] * r[i]
    }

    return res
}
