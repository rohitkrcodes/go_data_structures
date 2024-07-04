package main

import "github.com/rohitkrcodes/go_data_structures/lists"

func main() {
	ll := lists.NewDoublyList()

	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	ll.Append(4)
	ll.Append(5)

	ll.Prepend(0)

	ll.Remove(3)
	ll.PrintList()
}
