package collection

// LinkedList is an implementation of singly linked list.
// This is not concurrent safe.
type LinkedList[T any] struct {
	head   *linkedNode[T]
	tail   *linkedNode[T]
	length int
}

type linkedNode[T any] struct {
	v    T
	next *linkedNode[T]
}

// NewLinkedList returns an ArrayList based on the specified type.
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Add appends a given value to the bottom of the list.
// This is O(1) because LinkedList internally has the pointer to the tail node.
func (l *LinkedList[T]) Add(v T) {
	if l.head == nil {
		// if head is nil, the list has no elements.
		n := &linkedNode[T]{v: v, next: nil}
		l.head = n
		l.tail = n
		l.length++
		return
	}

	n := &linkedNode[T]{v: v, next: nil}
	l.tail.next = n
	l.tail = n
	l.length++
}

// AddHead inserts the given value at the head of the list.
func (l *LinkedList[T]) AddHead(v T) {
	if l.head == nil {
		// if head is nil, the list has no elements.
		n := &linkedNode[T]{v: v, next: nil}
		l.head = n
		l.tail = n
		l.length++
		return
	}

	curHead := l.head
	n := &linkedNode[T]{v: v, next: curHead}
	l.head = n
	l.length++
}

// AddAt appends a given value at the given index position in the list.
func (l *LinkedList[T]) AddAt(index int, v T) error {
	if index < 0 || l.length < index {
		return ErrInvalidIndex
	}

	if index == l.length {
		l.Add(v)
		return nil
	}

	if index == 0 {
		l.AddHead(v)
		return nil
	}

	n := &linkedNode[T]{v: v}

	var curr *linkedNode[T]
	for i := 0; i < index; i++ {
		if i == 0 {
			curr = l.head
		} else {
			curr = curr.next
		}
	}

	next := curr.next
	curr.next = n
	n.next = next
	l.length++
	return nil
}