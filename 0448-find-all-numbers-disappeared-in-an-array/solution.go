func findDisappearedNumbers(nums []int) []int {
    m := make(map[int]int, len(nums))

    for _, v :=  range nums {
        m[v]++
    }

    var msing []int
    for i, _ := range nums {
        if _, ok := m[i+1]; !ok {
            msing = append(msing, (i+1))
        }
    }
    return msing
}
