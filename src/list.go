package main

import "fmt"

type Nodo struct {
	data int
	next *Nodo
}

func newnodo(data int) *Nodo {
	return &Nodo{data, nil}
}

type List struct { //there isn't the only way, but for me it's the best
	size uint64
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
	list.size++
}

func (list *List) len() uint64 {
	return list.size
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
	list.size++
}

func (list *List) remove(data int) {

	if list.head == nil {
		return
	}

	current := list.head
	for current != nil {
		if current.next.data == data {
			current.next = current.next.next //that's funny!
			list.size--
			return
		}
		current = current.next
	}
}

func (list *List) display() {
	l := list.head
	for l != nil {
		fmt.Print("-> ", l.data)
		l = l.next
	}
}
