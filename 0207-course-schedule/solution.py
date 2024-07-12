class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        """
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
        """

        preMap : dict[int, list[int]] = {}
        for c in range(numCourses):
            preMap[c] = []
        for c, p in prerequisites:
            preMap[c].append(p)


        visited = set()

        def dfs(c: int) -> bool:
            if c in visited:
                return False
            if preMap[c] == []:
                return True
            
            visited.add(c)
            for p in preMap[c]:
                if not dfs(p):
                    return False
            visited.remove(c)
            preMap[c] = []

            return True
        
        for c in range(numCourses):
            if not dfs(c):
                return False


        return True
