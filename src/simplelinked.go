package list

import "fmt"

type Nodo struct {
	data int
	next *Nodo
}

func newnodo(data int) *Nodo {
	return &Nodo{data, nil}
}

type List struct { //there isn't the only way, but for me it's the best
	head *Nodo
}

func (list *List) append_front(data int) {

	new := newnodo(data)
	if list.head == nil {
		list.head = newnodo(data)
	} else {
		new.next = list.head
		list.head = new
	}
}

func (list *List) append(data int) {
	new := newnodo(data)
	if list.head == nil {
		list.head = newnodo(data)
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = new
	}
}

func (list *List) remove(data int) {

	current := list.head
	for current != nil {
		if current.data == data {
			current.next = current.next.next //that's funny!
			break
		}
		current = current.next
	}

}

func (list *List) display() {
	l := list.head
	for l != nil {
		fmt.Println(l.data)
		l = l.next
	}
}
