package benchmark

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	tail *node
}

func (l *linkedList) push(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}

	l.tail.next = n
	l.tail = n
}

func (l *linkedList) traverse(action func(int)) error {
	element := l.head
	for element != nil {
		action(element.data)
		element = element.next
	}

	return nil
}
