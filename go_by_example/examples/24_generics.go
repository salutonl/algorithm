package examples



func SliceIndex[E ~[]T, T comparable](anySlice E, value T) (int, T) {
	for i := range anySlice {
		if anySlice[i] == value {
			return i, anySlice[i]
		}
	}
	var empty T
	return -1, empty
}

type Node[T any] struct {
	value T
	next *Node[T]
}

type List[T any] struct {
	head, tail *Node[T]
}

func (l *List[T]) Push(v T) (T, bool) {
	if l.tail != nil {
		node := &Node[T]{value: v}
		l.tail.next = node
		l.tail = l.tail.next
	} else {
		l.head = &Node[T]{value: v}
		l.tail = l.head
	}
	return l.tail.value, true
}

func (l *List[T]) AllElements() []T {
	var ele []T

	for i := l.head; i != nil; i = i.next{
		ele = append(ele, i.value)
	}
	return ele
}