/*
	this is the file for implementing the doubly ended linked list
	this is implemented for a LRU cache use-case. it's not a general purpose linked list.
	V1:
		support only int values
		add new values to end
		don't allow add at a particular position
		delete the value by node
		don't allow delete via position
*/

package dell

import (
	"errors"
	"fmt"
)

type Node struct {
	key string
	value int
	prev *Node
	next *Node
}

type Dell struct {
	head *Node
	tail *Node
	length int
}

func New() *Dell {
	return &Dell {
		head: nil,
		tail: nil,
		length: 0,
	}
}

func (d *Dell) Add(key string, value int) (*Node, error) {
	node := &Node{key: key, value: value}

	if (d.length == 0) {
		d.head = node
		d.tail = node
		d.length++
		return node, nil
	}

	d.AddNodeToLast(node)
	return node, nil
}

func freeNode(node *Node) {
	// node will be garbage collected
	// we don't free memory manually like C/C++
	node.prev = nil
	node.next = nil
}

func (d *Dell) Delete(node *Node) (bool, error) {
	if (node == nil || d.length == 0) {
		return false, errors.New("invalid delete")
	}
	
	if (node == d.tail) {
		d.tail = node.prev
	} else if (node == d.head) {
		d.head = node.next
	} else {
		prev := node.prev
		next := node.next

		prev.next = next
		next.prev = prev
	}

	freeNode(node)
	d.length--
	return true, nil
}

// no one is using it externally
func (d *Dell) AddNodeToLast(node *Node) (bool, error) {
	node.prev = d.tail
	d.tail.next = node
	d.tail = node
	d.length++
	return true, nil
}

func (n *Node) UpdateValue(value int) (bool, error) {
	n.value = value
	return true, nil
}

func (n *Node) GetValue() int {
	return n.value
}

func (n *Node) GetKey() string {
	return n.key
}

func (d *Dell) GetLength() int {
	return d.length
}

func (d *Dell) MoveNodeToLast(node *Node) (bool, error) {
	if (node == d.tail) {
		return true, nil
	}
	d.Delete(node)
	d.AddNodeToLast(node)
	return true, nil
}

func (d *Dell) GetFirstNode() (*Node, error) {
	if (d.length == 0) {
		return nil, errors.New("list is empty")
	}

	return d.head, nil
}

func (d *Dell) GetLastNode() (*Node, error) {
	if (d.length == 0) {
		return nil, errors.New("list is empty")
	}

	return d.tail, nil
}

func (d *Dell) Print() {
	node := d.head
	fmt.Println("Printing list:")
	for node != nil {
		fmt.Printf("key=%q value=%d\n", node.key, node.value)
		node = node.next
	}
	fmt.Println("End of list")
}