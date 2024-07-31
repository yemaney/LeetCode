# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        """
        Given the root of a binary tree, return the values of the nodes
        you can see ordered from top to bottom when looking from the right side.
        
        Intuition:
        ----------
        - Use a breadth-first search (BFS) to traverse the tree level by level.
        - For each level, the last node encountered is the rightmost node visible
          from that level when viewed from the right side.
        
        Steps:
        -----
        1. Initialize a deque for BFS and append the root node to it.
        2. Initialize an empty list to store the right side view values.
        3. While the deque is not empty, repeat the following:
           - Determine the number of nodes at the current level (level_size).
           - Iterate through all nodes at the current level:
             - Pop the leftmost node from the deque.
             - Update the rightmost node's value for the current level.
             - Append the left and right children of the node to the deque if they exist.
           - Append the rightmost node's value to the right side view list.
        4. Return the list containing the right side view values.
        """
        if root is None:
            return []
        
        que : deque[TreeNode] = deque()
        que.append(root)
        ans : list[int] = []

        while que:
            level_right : int
            level_size : int = len(que)

            for _ in range(level_size):
                node = que.popleft()
                level_right = node.val

                if node.left:
                    que.append(node.left)
                if node.right:
                    que.append(node.right)
            ans.append(level_right)


        return ans
