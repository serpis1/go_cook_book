package main

import (
	"fmt"
	"sync"
)

type MyQueue struct {
	queue []string
	lock sync.RWMutex
}

func (q *MyQueue) Enqueue(name string) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.queue = append(q.queue, name)
}

func (q *MyQueue) Dequeue() error {
	if len(q.queue) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()
		q.queue = q.queue[1:]
		return nil
	}
	return fmt.Errorf("pop error: queue is empty")
}

func (q *MyQueue) Front() (string, error) {
	if len(q.queue) > 0 {
		q.lock.Lock()
		defer q.lock.Unlock()
		return q.queue[0], nil
	}
	return "", fmt.Errorf("pop error: queue is empty")
}

func (q *MyQueue) IsEmpty() bool {
	return len(q.queue) == 0
}

func (q *MyQueue) Size() int {
	return len(q.queue)
}

func (q *MyQueue) Print() {
	for val := range(q.queue) {
		fmt.Println(q.queue[val])
	}
}

func main() {

	myQueue := &MyQueue{
		queue: make([]string, 0),
	}

	myQueue.Enqueue("Bill")
	myQueue.Enqueue("Sara")
	myQueue.Enqueue("Max")
	myQueue.Enqueue("Peter")
	myQueue.Print()

	fmt.Println()

	fmt.Println(myQueue.Size())
	val, _ := myQueue.Front()
	fmt.Println(val)
	fmt.Println()

	myQueue.Dequeue()
	myQueue.Print()
	
}