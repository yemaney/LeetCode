// func missingNumber(nums []int) int {
//     max := 0
//     sum := 0
//     for i, v := range nums {
//         max += (i+1)
//         sum += v
//     }
//     return max - sum
// }
func missingNumber(nums []int) int {
    n := 0 
    for i, v := range nums {
        n ^= (i+1)
        n ^= v
    }
    return n
}


