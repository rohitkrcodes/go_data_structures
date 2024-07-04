package lists

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAppend(t *testing.T) {
	ll := NewDoublyList()

	for i := 0; i < 10; i++ {
		ll.Append(i)
	}

	require.Equal(t, 10, ll.Size)

	node := ll.Head
	for i := 0; i < 10; i++ {
		require.Equal(t, i, node.data)
		node = node.right
	}

	require.Nil(t, node)
}

func TestPrepend(t *testing.T) {
	ll := NewDoublyList()

	for i := 0; i < 10; i++ {
		ll.Prepend(i)
	}

	require.Equal(t, 10, ll.Size)

	node := ll.Tail
	for i := 0; i < 10; i++ {
		require.Equal(t, i, node.data)
		node = node.left
	}

	require.Nil(t, node)
}

// Popleft from an empty list
func TestPopleftFromEmptyList(t *testing.T) {
	ll := NewDoublyList()
	_, err := ll.Popleft()

	require.Error(t, err)
	require.EqualError(t, err, "list is empty")
}

// Popleft from an single element list
func TestPopleftSingleElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(10)
	res, err := ll.Popleft()

	require.Equal(t, 10, res)
	require.Nil(t, err)
	require.Nil(t, ll.Head)
	require.Nil(t, ll.Tail)
}

// Popleft from a two element list
func TestPopleftTwoElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)
	res, err := ll.Popleft()

	require.Nil(t, err)
	require.Equal(t, 1, res)
	require.Equal(t, 2, ll.Head.data)
	require.Equal(t, 2, ll.Tail.data)
}

// Popleft from a three element list
func TestPopleftThreeElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	res, err := ll.Popleft()

	require.Nil(t, err)
	require.Equal(t, 1, res)
	require.Equal(t, 2, ll.Head.data)
	require.Equal(t, 3, ll.Tail.data)
}

// Popright from an empty element list
func TestPoprightFromEmptyList(t *testing.T) {
	ll := NewDoublyList()
	_, err := ll.Popright()

	require.Error(t, err)
	require.EqualError(t, err, "list is empty")
}

// Popright from a single element list
func TestPoprightSingleElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(10)
	res, err := ll.Popright()

	require.Nil(t, err)
	require.Equal(t, 10, res)
	require.Nil(t, ll.Head)
	require.Nil(t, ll.Tail)
}

// Popright from a two element list
func TestPoprightTwoElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)
	res, err := ll.Popright()

	require.Nil(t, err)
	require.Equal(t, 2, res)
	require.Equal(t, 1, ll.Tail.data)
	require.Equal(t, 1, ll.Head.data)
}

// Popright from a three element list
func TestPoprightThreeElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)
	res, err := ll.Popright()

	require.Nil(t, err)
	require.Equal(t, 3, res)
	require.Equal(t, 2, ll.Tail.data)
	require.Equal(t, 1, ll.Head.data)
}

// Remove an item from an empty list
func TestRemoveFromEmptyList(t *testing.T) {
	ll := NewDoublyList()
	err := ll.Remove(10)

	require.Error(t, err)
	require.EqualError(t, err, "list is empty")
}

// Remove an item that does not exist in the list
func TestRemoveNonExistentElement(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	err := ll.Remove(10)

	require.Error(t, err)
	require.EqualError(t, err, "item not present in the list")
}

// Remove an existing item from a single element list
func TestRemoveSingleElementList(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	err := ll.Remove(1)

	require.Nil(t, err)
	require.Equal(t, 0, ll.Size)
	require.Nil(t, ll.Head)
	require.Nil(t, ll.Tail)
}

// Remove an existing item from the head of the list
func TestRemoveHead(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)

	err := ll.Remove(1)
	require.Nil(t, err)
	require.Equal(t, 1, ll.Size)
	require.Equal(t, 2, ll.Head.data)
	require.Nil(t, ll.Head.left)
}

// Remove an existing item from the tail of the list
func TestRemoveTail(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)

	err := ll.Remove(2)
	require.Nil(t, err)
	require.Equal(t, 1, ll.Size)
	require.Equal(t, 1, ll.Tail.data)
	require.Nil(t, ll.Tail.right)
}

// Remove an existing item from inside the list and not in the head or tail of the list
func TestRemoveMiddleElement(t *testing.T) {
	ll := NewDoublyList()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	err := ll.Remove(2)
	require.Nil(t, err)
	require.Equal(t, 2, ll.Size)
	require.Equal(t, ll.Tail, ll.Head.right)
	require.Equal(t, ll.Head, ll.Tail.left)
}
