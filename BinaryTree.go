package main

import (
	"fmt"
)

type BinaryTree struct {
	Val   int
	Left  *BinaryTree
	Right *BinaryTree
}

func Insert(v int, t *BinaryTree) *BinaryTree {
	if t == nil {
		return &BinaryTree{v, nil, nil}
	}
	if v < t.Val {
		t.Left = Insert(v, t.Left)
		return t
	}
	t.Right = Insert(v, t.Right)
	return t
}

func Find(v int, t *BinaryTree) *BinaryTree {
	if t == nil {
		return nil
	}
	if v < t.Left.Val {
		return Find(v, t.Left)
	} else if v > t.Right.Val {
		return Find(v, t.Right)
	}
	return t
}

func FindMin(t *BinaryTree) *BinaryTree {
	if t == nil {
		return nil
	}
	if t.Left == nil {
		return t
	}
	return FindMin(t.Left)
}

func FindMax(t *BinaryTree) *BinaryTree {
	if t == nil {
		return nil
	}
	if t.Right == nil {
		return t
	}
	return FindMax(t.Right)
}

func PrintTree(t *BinaryTree) {
	fmt.Println(t.Val)
	if t.Left != nil {
		PrintTree(t.Left)
	}
	if t.Right != nil {
		PrintTree(t.Right)
	}
}

func main() {
	t := Insert(100, nil)
	Insert(122, t)
	Insert(21, t)
	Insert(26, t)
	PrintTree(t)
	//	PrintTree(t)
}
