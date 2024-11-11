func construct2DArray(original []int, m int, n int) [][]int {
	/*
		Notes
		-----
		The function follows these steps:
		- Verify if the product `m * n` matches the length of `original`.
		- Return an empty list if the total number of elements does not match.
		- Initialize an empty list `ans` to store the final 2D array.
		- Use a loop to iterate through `original` and build subarrays of length `n`.
		- Append each completed subarray to `ans` and reset `subarray` for the next row.
		- Return `ans` after processing all elements.
	*/

	ans := [][]int{}
	if m*n != len(original) {
		return ans
	}

	i := 0
	subarray := []int{}
	for i < len(original) {
		subarray = append(subarray, original[i])
		i++

		if len(subarray) == n {
			ans = append(ans, subarray)
			subarray = []int{}
		}
	}

	return ans
}
