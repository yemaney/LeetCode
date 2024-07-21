# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrderBottom(self, root: Optional[TreeNode]) -> List[List[int]]:
        """
        Intuition:
        ----------
        - Use a queue to perform a level order traversal.
        - Insert each level's nodes at the beginning of the result list to achieve bottom-up order.

        Steps:
        ------
        1. Initialize an empty result list `res`.
        2. If the root is None, return the empty result list.
        3. Use a queue to perform level order traversal starting with the root.
        4. For each level:
            - Initialize an empty list `level` to store nodes of the current level.
            - Process all nodes at the current level.
            - Append the value of each node to the `level` list.
            - Add the left and right children of each node to the queue for the next level.
            - Insert the `level` list at the beginning of the result list `res`.
        5. Return the result list `res`.
        """

        res : list[list[int]] = []

        if root is None:
            return res

        que : deque[TreeNode] = deque()
        que.append(root)

        while len(que) > 0:
            level : list[int] = []

            for _ in range(len(que)):
                node = que.popleft()

                level.append(node.val)

                if node.left:
                    que.append(node.left)
                if node.right:
                    que.append(node.right)

            res.insert(0, level)


        return res
