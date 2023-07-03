func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    m := make(map[int]bool, len(nums))
    for _, v := range nums {
        m[v] = false
    }


    maxSeq := 0
    for _, v := range nums {
        s := 1

        for p := v; ; p++ {
            if pop, ok := m[p+1]; ok {
                if pop {
                    break
                }
                s++
                m[p+1] = true
            } else {
                break
            } 
        }
        for p := v; ; p-- {
            if pop, ok := m[p-1]; ok {
                if pop {
                    break
                }
                s++
                m[p-1] = true
            } else {
                break
            } 
        }
        
        if s > maxSeq {
            maxSeq = s
        }
    }

    return maxSeq
}
