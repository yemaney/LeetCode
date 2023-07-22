import  "fmt"
func sortColors(nums []int)  {
    l := 0
    r := len(nums) - 1

    for l < r {
        if nums[l] == 0 {
            l++
        } else if nums[r] == 0 {

            nums[l], nums[r] = nums[r], nums[l]
            l++
            r--
        } else {
            r--
        }
    }

    r = len(nums) - 1
    fmt.Println(l, r)
    for l < r {
        if nums[l] == 1 || nums[l] == 0{
            l++
        } else if nums[r] == 1 {

            nums[l], nums[r] = nums[r], nums[l]
            l++
            r--
        } else {
            r--
        }
    }
}
