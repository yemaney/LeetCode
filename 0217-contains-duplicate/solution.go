func containsDuplicate(nums []int) bool {
    cache := make(map[int]int, len(nums))
    for _, v := range nums {
        if _, ok := cache[v]; ok {
            return true
        }
        cache[v] = 0
    }
    return false
}
