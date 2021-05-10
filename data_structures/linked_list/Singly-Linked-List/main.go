package main

import (
	"fmt"
	"sync"
)

// Node a single node that composes the list
type Node struct {
	content string
	next *Node
}

type ItemLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

func (i *ItemLinkedList) Append(s string) {
	i.lock.Lock()
	node := Node{
		content: s,
		next:    nil,
	}

	if i.head == nil {
		i.head = &node
	} else {
		last := i.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	i.size++
	i.lock.Unlock()
}


func (i *ItemLinkedList) Insert(pos int, s string) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	if pos < 0 || pos > i.size {
		return fmt.Errorf("index out of range")
	}

	node := Node{
		content: s,
		next:    nil,
	}
	if pos == 0 {
		node.next = i.head
		i.head = &node
		return nil
	}
	curNode := i.head
	j := 0

	for j < pos - 2 {
		j++
		curNode = curNode.next
	}
	node.next = curNode.next
	curNode.next = &node
	i.size++
	return nil
}

func (i *ItemLinkedList) RemoveAt(pos int) (*string, error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	if pos < 0 || pos > i.size {
		return nil, fmt.Errorf("Out of range")
	}

	node := i.head
	j := 0
	for j < pos - 1 {
		j++
		node = node.next
	}
	remove := node.next
	node.next = remove.next
	i.size--

	return &remove.content, nil
}


func (i *ItemLinkedList) IndexOf(s string) int {
	i.lock.Lock()
	defer i.lock.Unlock()

	node := i.head
	j := 0
	for {
		if node.content == s {
			return j
		}
		if node.next == nil {
			return -1
		}
		node = node.next
		j++
	}
}


func (i *ItemLinkedList) IsEmpty() bool {
	i.lock.Lock()
	defer i.lock.Unlock()

	if i.head == nil {
		return true
	}
	return false
}


func (i *ItemLinkedList) Size() int {
	return i.size
}

func (i *ItemLinkedList) Reverse() {
	i.lock.Lock()
	defer i.lock.Unlock()

	var cursor *Node = i.head
	var prev *Node = nil

	for cursor != nil {
		next := cursor.next
		cursor.next = prev
		prev = cursor
		cursor = next
	}
	i.head = prev
}


func (i *ItemLinkedList) Print() {
	i.lock.Lock()
	defer i.lock.Unlock()

	node := i.head

	for {
		if node == nil {
			break
		}

		fmt.Print(node.content)
		fmt.Print(" -> ")
		node = node.next
	}
	fmt.Println("")
}


func main() {
	sLinkedList := &ItemLinkedList{
		head: nil,
		size: 0,
		lock: sync.RWMutex{},
	}
	fmt.Println(sLinkedList.IsEmpty())
	sLinkedList.Append("Pablo")
	sLinkedList.Append("Anna")
	sLinkedList.Append("Antonio")
	sLinkedList.Insert(1, "Serhio")

	fmt.Println(sLinkedList.IsEmpty())
	fmt.Println(sLinkedList.IndexOf("Anna"))

	sLinkedList.RemoveAt(0)
	fmt.Println(sLinkedList.Size())
	sLinkedList.Print()
	sLinkedList.Reverse()
	sLinkedList.Print()
}
