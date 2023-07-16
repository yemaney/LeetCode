import (
	"sort"
)

func threeSum(nums []int) [][]int {
    var result [][]int

    sort.Ints(nums)

    for i := 0; i <= len(nums) - 3; i++ {
        result = bsearch(result, nums, i)
    }

    return result
}

func bsearch(result [][]int, nums []int, start int) [][]int {
    if start > 0 && nums[start] == nums[start - 1] {
        return result
    }

    l := start + 1
    r := len(nums) - 1

    want := 0 - nums[start]
    for l < r {
        if nums[l] + nums[r] == want {
            result = append(result, []int{nums[start], nums[l], nums[r]})
            l = movePointer(nums, l, true)
            r = movePointer(nums, r, false)
        } else if nums[l] + nums[r] < want {
            l = movePointer(nums, l, true)
        } else if nums[l] + nums[r] > want {
            r = movePointer(nums, r, false)
        }
    }

    return result
}

func movePointer(nums []int, i int, increment bool) int {
    if increment {
        j := i + 1
        for j < len(nums) && nums[i] == nums[j] {
            j++
        }
        return j
    }
    j := i - 1
    for j > 0 && nums[i] == nums[j] {
        j--
    }
    return j
}

