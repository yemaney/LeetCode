func twoSum(nums []int, target int) []int {
	cache := map[int]int{}
	ret := []int{}
out:
	for k, v := range nums {
		compliment := target - v

		if pop, ok := cache[compliment]; ok {
			ret = append(ret, k, pop)
			break out

		} else {
			cache[v] = k
		}

	}
	return ret  
}
