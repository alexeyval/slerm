package main

func main() {
	list := createList([]int{1, 2, 3, 3})
	noDuplicates := deleteDuplicates(list)
	_ = noDuplicates
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// createList - создаёт цепочку структур, связный список
//
//	slice - отсортированный по возрастанию
func createList(slice []int) *ListNode {
	if len(slice) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  slice[0],
		Next: nil,
	}
	node := head

	for _, v := range slice[1:] {
		node.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		node = node.Next
	}

	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	node := head
	for node != nil {
		for node.Next != nil && node.Val == node.Next.Val {
			node.Next = node.Next.Next
		}
		node = node.Next
	}
	return head
}
