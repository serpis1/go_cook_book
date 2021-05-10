package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MaxLevel = 32
	P = 1 / math.E
)

type Node struct {
	score int
	next []*Node
}

func nodeCreator(score int, level int) *Node {
	return &Node{
		score: score,
		next:  make([]*Node, level),
	}
}

type Skiplist struct {
	head *Node
	level int
}

func Constructor() Skiplist {
	return Skiplist{
		head:  nodeCreator(0, 32),
		level: 1,
	}
}

func randomLevel() int {
	level := 1
	for rand.Float64() < P {
		level++
	}
	if level > MaxLevel {
		return MaxLevel
	}
	return level
}

func (s *Skiplist) Search(value int) bool {
	x := s.head
	for i := s.level - 1; i >= 0; i-- {
		for x.next[i] != nil {
			if x.next[i].score == value {
				return true
			} else if x.next[i].score < value {
				x = x.next[i]
			} else {
				break
			}
		}
	}
	return false
}

func (s *Skiplist) Add(value int) {
	update := make([]*Node, MaxLevel)

	for i, x := s.level-1, s.head; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].score < value {
			x = x.next[i]
		}
		update[i] = x
	}

	level := randomLevel()
	n := nodeCreator(value, level)
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}

	for i := 0; i < level; i++ {
		n.next[i] = update[i].next[i]
		update[i].next[i] = n
	}
}

func (s *Skiplist) Erase(val int) bool {
	update := make([]*Node, MaxLevel)
	for i, x := s.level-1, s.head; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].score < val {
			x = x.next[i]
		}
		update[i] = x
	}

	node := update[0].next[0]
	if node == nil || node.score != val {
		return false
	}

	for i := 0; i < len(node.next); i++ {
		update[i].next[i] = node.next[i]
	}

	for s.level > 1 && s.head.next[s.level-1] == nil {
		s.level--
	}
	return true
}

func (s *Skiplist) draw() {
	fmt.Println('\n')

	for level := s.level; level >= 0; level-- {
		if s.head.next[level] == nil { continue }
		fmt.Print(level, ": ")
		for node := s.head.next[level]; node != nil; node = node.next[level] {
			fmt.Print("--", node.score, "--")
		}
		fmt.Println('\n')
	}
	fmt.Println('\n')
}

func main()  {
	skipLst := Constructor()
	for i := 0; i < 30; i++ {
		skipLst.Add(i)
	}
	skipLst.Erase(6)
	skipLst.draw()
	//fmt.Println('\n', MaxLevel, P, skipLst)
	fmt.Println(skipLst.Search(4))

}

