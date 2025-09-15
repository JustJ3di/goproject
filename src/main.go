package main

import "fmt"

func main() {

	mylist := new(List)

	for i := 0; i < 10; i++ {
		mylist.append(i)
	}

	fmt.Println("before", mylist.len())

	mylist.remove(4)

	fmt.Println(mylist.len())

	mylist.display()

}
