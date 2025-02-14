package linkedlist

type Node struct {
	data Student
	next *Node

}

func LinkedListToString(head *Node) (str string) {
	for cur := head; cur != nil; cur = cur.next {
		str += " -> " + StudentToString(cur.data)
	}

	return str
}

func LinkedListAdd(head **Node, student Student){
	// node := &Node {
	// 	data: student,
	// }

	node := new(Node)
	node.data = student 

	if *head == nil {
		*head = node
	}

	tail := *head
	for tail.next != nil {
		tail = tail.next
	}

	tail.next = node
}