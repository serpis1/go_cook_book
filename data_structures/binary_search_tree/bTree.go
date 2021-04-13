package main

import (
	"fmt"
)


type TreeNode struct {
	val int
	left *TreeNode
	right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return fmt.Errorf("Tree is nil.")
	}

	if t.val == value {
		return fmt.Errorf("This node value is already exist")
	}

	if t.val > value {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return nil
		}
		return t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return nil
		}
		return t.right.Insert(value)
	}
	return nil
}


func (t *TreeNode) PrintTree() {
	if t == nil {
		return
	}

	t.left.PrintTree()
	fmt.Println(t.val)
	t.right.PrintTree()
}

func (t *TreeNode) Min() int {
	if t.left == nil {
		return t.val
	}
	return t.left.Min()
}

func (t *TreeNode) Max() int {
	if t.right == nil {
		return t.val
	}
	return t.right.Max()
}


func main() {
	t := &TreeNode{val: 8}
	t.Insert(2)
	t.Insert(1)
	t.Insert(5)
	t.Insert(56)

	t.PrintTree()
	fmt.Println(t.Min(), t.Max())
}