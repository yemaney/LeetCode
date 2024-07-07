# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def averageOfLevels(self, root: Optional[TreeNode]) -> List[float]:
        ans : List[float] = []
        if root is None:
            return ans
        
        q : List[TreeNode] = []
        q.append(root)

        while len(q) > 0:
            size = len(q)
            sum = 0 
            
            for _ in range(size):
                node = q[0]
                q = q[1:]

                sum += node.val

                if node.left is not None:
                    q.append(node.left)
                if node.right is not None:
                    q.append(node.right)            

            avg = sum / size
            ans.append(avg)

        return ans
