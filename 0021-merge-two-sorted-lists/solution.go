/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 == nil {
        return nil
    }
    x := &ListNode{}
    p := x
    for list1 != nil || list2 != nil {
        // finished reading list1
        if list1 == nil {
            p.Val = list2.Val
            p.Next = list2.Next
            break
        // finished reading list2
        } else if list2 == nil {
            p.Val = list1.Val
            p.Next = list1.Next
            break
        // current node at list1 is lower
        } else if list1.Val < list2.Val{
            p.Val = list1.Val
            list1 = list1.Next
            p.Next = &ListNode{}
            p = p.Next
        // current node at list2 is lower
        } else {
            p.Val = list2.Val
            list2 = list2.Next
            p.Next = &ListNode{}
            p = p.Next
        }
    }
    
    return x
}
