
func findOrder(numCourses int, prerequisites [][]int) []int {
	/*
	   Determines the order in which courses should be taken given the prerequisites.

	   Args:
	       numCourses (int): The total number of courses.
	       prerequisites (List[List[int]]): A list where each element is a pair [a, b] indicating that course 'a' has a prerequisite course 'b'.

	   Returns:
	       List[int]: A list of courses in the order they should be taken to complete all courses.
	                 If it is impossible to complete all courses (due to a cycle), returns an empty list.

	   Steps:
	       1. Create a map (`preMap`) to store course and its prerequisite courses as key-value pairs.
	       2. For each course, perform a depth-first search (DFS) to traverse the prerequisite chain.
	       3. Use a set `visited` to track nodes that have been fully processed.
	       4. Use a set `curCycle` to track nodes currently being visited in the current DFS cycle to detect cycles.
	       5. If a node is visited twice in the same cycle, a cycle is detected, and the function returns an empty list.
	       6. If a course with no prerequisites is found, it is appended to the result list `ans`.
	       7. If a course with every prerequisite already processed is found, it is appended to the result list `ans`.

	   Example:
	       prerequisites = [[1,0], [0,2], [1,2]]

	       The function should return a list such as [2, 0, 1] or [2, 1, 0].

	       Explanation:
	           preMap = {
	               0: [2],
	               1: [0, 2],
	               2: []
	           }
	*/
	ans := []int{}

	preMap := make(map[int][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		preMap[prerequisites[i][0]] = append(preMap[prerequisites[i][0]], prerequisites[i][1])
	}

	visited := make(map[int]bool, numCourses)
	curCycle := make(map[int]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if dfs(i, &ans, preMap, visited, curCycle) == false {
			return []int{}
		}
	}

	return ans
}

func dfs(c int, ans *[]int, preMap map[int][]int, visited map[int]bool, curCycle map[int]bool) bool {
	if _, ok := curCycle[c]; ok {
		return false
	}
	if _, ok := visited[c]; ok {
		return true
	}

	curCycle[c] = true
	// The first time this passes as true, is when `c` doesn't have any preqs.
	// All passes occur if c has no preqs or if they have all been already been processed.
	for _, v := range preMap[c] {
		if dfs(v, ans, preMap, visited, curCycle) == false {
			return false
		}
	}

	delete(curCycle, c)
	visited[c] = true
	*ans = append(*ans, c)

	return true
}

