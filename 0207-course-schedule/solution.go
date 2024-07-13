func canFinish(numCourses int, prerequisites [][]int) bool {
	/*
		1. create map of course : preqs pairs
		2. for each course do dfs to traverse the preq chain
		3. store visited nodes in set along the way
		4. if a node is visited twice exit early as False, else True

		This can be done because of the property `0 <= ai, bi < numCourses`.
		The values for courses and preqs being within the bounds of numCourses
		enables hopping through preMap as a tree.

		Example:
			prerequisites = [[1,0], [0,2], [1,2]]

			preMap = {
				0 : [2],
				1 : [0, 2],
				2 : []
			}
	*/
	preMap := map[int][]int{}
	for i := 0; i < len(prerequisites); i++ {
		c, p := prerequisites[i][0], prerequisites[i][1]
		preMap[c] = append(preMap[c], p)
	}

	visited := map[int]bool{}

	for i := 0; i < numCourses; i++ {
		if dfs(i, preMap, visited) == false {
			return false
		}
	}

	return true
}

func dfs(x int, preMap map[int][]int, visited map[int]bool) bool {
	if _, ok := visited[x]; ok {
		return false
	}
	if pop, ok := preMap[x]; ok && len(pop) == 0 {
		return true
	}

	visited[x] = true
	for _, v := range preMap[x] {
		if dfs(v, preMap, visited) == false {
			return false
		}
	}

	delete(visited, x)
	preMap[x] = []int{}

	return true
}

