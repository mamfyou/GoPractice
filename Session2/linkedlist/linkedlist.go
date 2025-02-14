package linkedlist

type Node struct {
	Value int
	Next  *Node
}

func MergeNodes(head *Node) *Node {
	var head_node *Node = nil
	var current_node *Node = &Node{}
	if head == nil {
		return nil
	}

	if head.Next == nil {
		if head.Value != 0 {
			new_node := new(Node)
			new_node.Value = head.Value
			return new_node
		} else {
			return nil
		}
	}

	sum_numbers := 0
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Value == 0 {
			if sum_numbers > 0 {
				new_node := new(Node)
				new_node.Value = sum_numbers
				current_node.Next = new_node
				if head_node == nil {
					head_node = new_node
				}
				sum_numbers = 0
				current_node = current_node.Next
			}
		} else {
			sum_numbers += curr.Value
			if curr.Next == nil {
				new_node := new(Node)
				new_node.Value = sum_numbers
				current_node.Next = new_node

				if head_node == nil {
					head_node = new_node
				}
				current_node = current_node.Next
			}
		}
	}
	
	return head_node
}
