package algos

import "dsa/ds"

func MiddleOfLinkedList(head *ds.Node) int {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow.Data
}
