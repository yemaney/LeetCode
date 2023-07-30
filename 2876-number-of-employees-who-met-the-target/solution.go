func numberOfEmployeesWhoMetTarget(hours []int, target int) int {
    res := 0
    for _, v := range hours {
        if v >= target {
            res++
        }
    }
    return res
}
