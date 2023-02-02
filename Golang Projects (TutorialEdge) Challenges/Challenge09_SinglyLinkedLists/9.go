package Challenge09_SinglyLinkedLists

type Element struct {
	Value string
	Next  *Element
}

type SinglyLinkedList struct {
	Count int
	Head  *Element
	Tail  *Element
}

func (l *SinglyLinkedList) Append(value string) {
	element := &Element{Value: value}

	if l.Head == nil {
		l.Head = element
	} else {
		l.Tail.Next = element
	}

	l.Tail = element
	l.Count++
}

func (l *SinglyLinkedList) Size() int {
	return l.Count
}

func (l *SinglyLinkedList) Print() []string {
	var liste []string
	current := l.Head

	for current != nil {
		liste = append(liste, current.Value)
		current = current.Next
	}

	return liste
}

func Test() []string {
	var llist SinglyLinkedList

	values := []string{"First", "Second", "Third"}
	for _, value := range values {
		llist.Append(value)
	}

	return llist.Print()
}
