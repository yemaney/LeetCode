# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def minDepth(self, root: Optional[TreeNode]) -> int:
        """
        1. traverse tree by level
        2. increment the ans each level
        3. exit as soon as a root leaf is found        
        """

        if root is None:
            return 0

        ans : int = 0
        q : List[TreeNode] = []
        q.append(root)

        while q:
            ans += 1
            size : int = len(q)

            for _ in range(size):
                
                node : TreeNode = q[0]
                q = q[1:]

                if node.left is None and node.right is None:
                    return ans
                
                if node.left is not None:
                    q.append(node.left)
                if node.right is not None:
                    q.append(node.right)


        return ans
