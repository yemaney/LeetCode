class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        """
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

        Example:
            prerequisites = [[1,0], [0,2], [1,2]]
            
            The function should return a list such as [2, 0, 1] or [2, 1, 0].

            Explanation:
                preMap = {
                    0: [2],
                    1: [0, 2],
                    2: []
                }
        """
        preMap : dict[int, list[int]] = {i:[] for i in range(numCourses)}
        for c, p in prerequisites:
            preMap[c].append(p)
        

        ans : list[int] = []
        visited = set()
        curCycle = set()

        def dfs(c: int) -> bool:
            """
            Depth-First Search to detect cycles and determine the order of courses.
            The first time this passes as true, is when `c` doesn't have any preqs.
            All passes occur if c has no preqs or if its already been processed.

            Args:
                c (int): The current course being visited.

            Returns:
                bool: True if no cycle is detected, False if a cycle is detected.
            """
            if c in curCycle:
                return False
            if c in visited:
                return True
            
            curCycle.add(c)

            # The first time this passes as true, is when `c` doesn't have any preqs.
            # All passes occur if c has no preqs or if its already been processed.
            for p in preMap[c]:
                if not dfs(p):
                    return False
            
            curCycle.remove(c)
            visited.add(c)
            ans.append(c)

            return True
        
        for c in range(numCourses):
            if not dfs(c):
                return []
        
        return ans
