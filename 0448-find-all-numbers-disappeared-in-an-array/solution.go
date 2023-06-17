// func findDisappearedNumbers(nums []int) []int {
 
//     // map to hold found nums
//     m := make(map[int]int, len(nums))

//     // add each num in nums to map
//     for _, v :=  range nums {
//         m[v]++
//     }

//     // create a slice of all nums not found in the map 
//     var msing []int
//     for i, _ := range nums {
//         if _, ok := m[i+1]; !ok {
//             msing = append(msing, (i+1))
//         }
    
//     return msing
// }

func findDisappearedNumbers(nums []int) []int {
    missing := make([]int, len(nums))

    // using result array as sorted array for seen values
    for _, v := range nums {
        missing[v-1] = v
    }

    // here we have res array where missing records will be with value = 0
    // scan all values with 0 and "shift" them to head of res array
    j := 0
    for i, v := range missing {
        if v == 0 {
            missing[j] = i + 1
            j++
        }
    }
 
    // chop head of res array and return it
    return missing[:j]
}



