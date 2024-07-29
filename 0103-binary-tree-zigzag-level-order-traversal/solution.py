# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        """
        Perform a level order traversal on a binary tree and reverse the order of nodes at every alternate level.

        Intuition:
        The goal is to traverse a binary tree in level order, but with a twist: we want to reverse the order of the levels. 
        To achieve this, we use a queue to facilitate level-order traversal and a flag (`rev`) to determine whether to reverse 
        the current level's nodes before adding them to the result list.

        Steps:
        1. Initialization:
            - Initialize a flag `rev` to `False`. This flag will help decide whether to reverse the current level's nodes.
            - Initialize an empty list `ans` to store the final result.
            - If the root is `None`, return `ans` immediately as the tree is empty.
            - Initialize a deque `que` and append the root node to it.
        2. Level Order Traversal:
            - While the queue is not empty:
                1. Initialize an empty list `level` to store nodes of the current level.
                2. Iterate over the current size of the queue:
                    - Pop a node from the left side of the queue.
                    - Append its value to the `level` list.
                    - If the node has a left child, append it to the queue.
                    - If the node has a right child, append it to the queue.
                3. If the `rev` flag is `True`, reverse the `level` list.
                4. Append the `level` list to `ans`.
                5. Toggle the `rev` flag for the next level.
        3. Return the Result:
            - Return the `ans` list containing the nodes in the required order.
        """
        rev : bool = False 
        ans : list[list[int]] = []

        if not root:
            return ans
        
        que : deque[int] = deque()
        que.append(root)

        
        while que:
            level : list[int] = []

            for _ in range(len(que)):
                node : TreeNode = que.popleft()

                level.append(node.val)

                if node.left:
                    que.append(node.left)
                if node.right:
                    que.append(node.right)


            if rev:
                level.reverse()
            ans.append(level)
            rev = not rev        


        return ans
