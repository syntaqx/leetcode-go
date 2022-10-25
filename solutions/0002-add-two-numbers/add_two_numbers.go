package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := 0
	head := new(*ListNode)
	current := head
	for l1 != nil || l2 != nil || t != 0 {
		if l1 != nil {
			t += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			t += l2.Val
			l2 = l2.Next
		}
		if *current == nil {
			*current = new(ListNode)
		}
		val := t % 10
		t /= 10
		(**current).Val = val
		current = &((*current).Next)
	}
	return *head
}
