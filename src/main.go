package main

func main() {

	mylist := new(List)

	for i := 0; i < 10; i++ {
		mylist.append_front(i)
	}

	mylist.append(12)

	mylist.remove(4)

	mylist.display()

}
