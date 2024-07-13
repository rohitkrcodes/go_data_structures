package lists

import (
	"fmt"
	"strconv"
)

type doublyNode struct {
	data  int
	left  *doublyNode
	right *doublyNode
}

func newDoublyNode(data int) *doublyNode {
	return &doublyNode{
		data:  data,
		left:  nil,
		right: nil,
	}
}

type DoublyList struct {
	Head *doublyNode
	Tail *doublyNode
	Size int
}

func NewDoublyList() *DoublyList {
	return &DoublyList{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// adds data item to the tail of the list
func (ll *DoublyList) Append(data int) error {

	// exceed max size
	if ll.Size > 1000000 {
		return fmt.Errorf("reached max size limit")
	}

	// create new node
	node := newDoublyNode(data)

	if ll.Size == 0 { // if list is empty
		ll.Head = node
		ll.Tail = node
	} else { // if list is non-empty
		ll.Tail.right = node
		node.left = ll.Tail
		ll.Tail = node
	}

	// increment list size by one
	ll.Size += 1

	return nil
}

// adds data item at the head of the list
func (ll *DoublyList) Prepend(data int) error {

	// exceed max size
	if ll.Size > 1000000 {
		return fmt.Errorf("reached max size limit")
	}

	// create new node
	node := newDoublyNode(data)

	if ll.Size == 0 { // if list is empty
		ll.Head = node
		ll.Tail = node
	} else { // if list is non-empty
		node.right = ll.Head
		ll.Head.left = node
		ll.Head = node
	}

	// increment list size by one
	ll.Size += 1

	return nil
}

// removes item from the left of list
func (ll *DoublyList) Popleft() (int, error) {

	data := 0
	if ll.Size == 0 { // Empty List
		return data, fmt.Errorf("list is empty")
	} else if ll.Size == 1 { // Only one element in list
		data = ll.Head.data
		ll.Head = nil
		ll.Tail = nil
	} else { // More than one element in list
		temp := ll.Head
		data = temp.data
		ll.Head = temp.right
		ll.Head.left = nil
		temp.right = nil
		temp = nil
	}

	// decrement list size by one
	ll.Size -= 1
	return data, nil
}

// removes item from the right of the list
func (ll *DoublyList) Popright() (int, error) {

	data := 0

	if ll.Size == 0 { // Empty List
		return data, fmt.Errorf("list is empty")
	} else if ll.Size == 1 { // Only one element in list
		data = ll.Tail.data
		ll.Head = nil
		ll.Tail = nil
	} else { // More than one element in list
		temp := ll.Tail
		data = temp.data
		ll.Tail = temp.left
		ll.Tail.right = nil
		temp.left = nil
		temp = nil
	}

	// decrement list size by one
	ll.Size -= 1
	return data, nil
}

// removes an existing data item from anywhere in the list
func (ll *DoublyList) Remove(data int) error {

	// empty list
	if ll.Size == 0 {
		return fmt.Errorf("list is empty")
	}

	node := ll.Head

	// loop through list until we get the element node to remove
	for node != nil {
		if node.data == data {
			break
		} else {
			node = node.right
		}
	}

	// if node is nil - implies element node to remove not found in the list
	if node == nil {
		return fmt.Errorf("item not present in the list")
	}

	// if the node with the element to remove is the head of the list
	if node.left == nil {

		// if only one element present in the list
		if ll.Size == 1 {
			ll.Head = nil
			ll.Tail = nil
			node = nil
			ll.Size -= 1
			return nil
		}

		// if size of list is greater than one
		temp := node.right
		temp.left = nil
		node.right = nil
		node = nil
		ll.Head = temp
		ll.Size -= 1
		return nil
	}

	// if the node with the element to remove is the tail of the list
	if node.right == nil {
		temp := node.left
		temp.right = nil
		node.left = nil
		node = nil
		ll.Tail = temp
		ll.Size -= 1
		return nil
	}

	// if the node with the element to remove is neither the head nor the tail of list
	prev := node.left
	next := node.right
	prev.right = next
	next.left = prev
	node.left = nil
	node.right = nil
	node = nil
	ll.Size -= 1
	return nil

}

func (ll *DoublyList) PrintList() {
	if ll.Size == 0 {
		fmt.Println("Empty list")
	}

	node := ll.Head
	fmt.Println("Size of list is", ll.Size)
	for node != nil {
		fmt.Println(node.data)
		node = node.right
	}
}

func (ll *DoublyList) GetList(from int, length int) (map[string]int, error) {
	if ll.Size == 0 {
		return nil, fmt.Errorf("empty list")
	}

	res := make(map[string]int)

	node := ll.Head
	i := 1
	for node != nil {

		if i >= from && i-from < length {
			x := fmt.Sprintf("data position %s", strconv.Itoa(i))
			res[x] = node.data
		}

		node = node.right
		i += 1
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("no data")
	}

	return res, nil
}

func (ll *DoublyList) GetData(data int) (map[string]int, error) {

	if ll.Size == 0 {
		return nil, fmt.Errorf("empty list")
	}

	res := make(map[string]int)

	node := ll.Head
	i := 1
	for node != nil {
		if node.data == data {
			x := fmt.Sprintf("data position %s", strconv.Itoa(i))
			res[x] = node.data
		}
		node = node.right
		i += 1
	}

	if len(res) == 0 {
		return nil, fmt.Errorf("data not found")
	}

	return res, nil

}
