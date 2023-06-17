func construct2DArray(original []int, m int, n int) [][]int {
    length := len(original)
    new := make([][]int, 0, m)

    // early exit if original array is to long to fit into an mxn 2d array 
    if length != m * n {
        return [][]int{}
    }

    // append each row 
    for i := 0; i < length; i+=n {
        new = append(new, original[i:i+n])
    }
    return new

}
