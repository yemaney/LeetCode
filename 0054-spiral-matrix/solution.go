func spiralOrder(matrix [][]int) []int {
    arr := make([]int, len(matrix) * len(matrix[0]))

    i := 0
    l, r, t, b := 0, len(matrix[0]), 0, len(matrix)
	outerLoop:
	for ; l < r &&  t < b; {
	    for j := l; j < r; j++ {
			if i == len(matrix) * len(matrix[0]) {
				l = r
				break outerLoop
			}
	        arr[i] = matrix[t][j]
	        i++
	    }
	    t++
	    for j := t; j < b; j++ {
			if i == (len(matrix) * len(matrix[0])) {
				l = r
				break outerLoop
			}
	        arr[i] = matrix[j][r-1]
	        i++
	    }
	    r--
	    for j := r ; j > l; j-- {
			if i == (len(matrix) * len(matrix[0])) {
				l = r
				break outerLoop
			}
	        arr[i] = matrix[b-1][j-1]
	        i++
	    }
		b--
	    for j := b ; j > t; j-- {
			if i == (len(matrix) * len(matrix[0])) {
				l = r
				break outerLoop
			}
	        arr[i] = matrix[j-1][l]
	        i++
	    }
		l++
		fmt.Println(i)
	}
    return arr
}
