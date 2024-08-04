# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def distanceK(self, root: TreeNode, target: TreeNode, k: int) -> List[int]:
        """
        Finds all nodes at distance K from the target node in a binary tree.

        Intuition:
        - Convert the binary tree into an adjacency list to handle it like an undirected graph.
        - Use Breadth-First Search (BFS) starting from the target node to find all nodes at distance K.

        Steps:
        1. Check if the root is None. If so, return an empty list.
        2. Initialize the adjacency list using the binary tree.
        3. Initialize a queue with the target node's value and a visited set containing the target node's value.
        4. Perform BFS until the distance equals K:
            - For each node at the current level, add its neighbors to the queue if they haven't been visited.
            - Increment the steps counter.
        5. Return the list of nodes in the queue.

        Args:
        root (TreeNode): The root of the binary tree.
        target (TreeNode): The target node.
        k (int): The distance from the target node.

        Returns:
        List[int]: A list of node values at distance K from the target node.
        """

        if not root:
            return []

        adj : dict[int, list[int]] = {root.val: []}
        init_adj(root, adj)


        que : deque[int] = deque([target.val])
        visited : set[int] = set([target.val])
        steps : int = 0

        while que:
            if steps == k:
                break

            q_size : int = len(que)

            for _ in range(q_size):
                node : int = que.popleft()
                next_level : list[int] = adj[node]
                
                for n in next_level:
                    if n not in visited:
                        que.append(n)
                        visited.add(n)

            steps += 1

        return list(que)


def init_adj(node: TreeNode, adj: dict[int, list[int]]) -> None:
    """
    Initializes the adjacency list for the binary tree.

    Intuition:
    - Traverse the binary tree and create bidirectional edges in the adjacency list for each node.

    Steps:
    1. If the node has a left child, add bidirectional edges between the node and the left child.
    2. Recursively initialize the adjacency list for the left child.
    3. If the node has a right child, add bidirectional edges between the node and the right child.
    4. Recursively initialize the adjacency list for the right child.

    Args:
    node (TreeNode): The current node in the binary tree.
    adj (dict[int, list[int]]): The adjacency list representing the binary tree as an undirected graph.

    Returns:
    None
    """
    if node.left:
        adj[node.val].append(node.left.val)
        adj[node.left.val] = [node.val]
        
        init_adj(node.left, adj)

    if node.right:
        adj[node.val].append(node.right.val)
        adj[node.right.val] = [node.val]

        init_adj(node.right, adj)


