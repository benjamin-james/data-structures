package main

import "fmt"

type node struct {
	value       element
	left, right *node
	height      int
}

func (n *node) init(value element) *node {
	n.left = nil
	n.right = nil
	n.height = 1
	n.value = value
	return n
}

func NewNode(value element) *node {
	return new(node).init(value)
}

func (n *node) Display() {
	if n.left != nil {
		n.left.Display()
	}
	fmt.Println(n.value)
	if n.right != nil {
		n.right.Display()
	}
}
func (n *node) diff() int {
	return n.left_height() - n.right_height()
}

func (n *node) left_height() int {
	if n.left != nil {
		return n.left.height
	} else {
		return 0
	}
}

func (n *node) right_height() int {
	if n.right != nil {
		return n.right.height
	} else {
		return 0
	}
}

func (n *node) get_height() int {
	lh, rh := n.left_height(), n.right_height()
	if lh > rh {
		n.height = 1 + lh
	} else {
		n.height = 1 + rh
	}
	return n.height
}

func (n *node) find(key element) *node {
	ret := key.Compare(n.value)
	if ret > 0 && n.right != nil {
		return n.right.find(key)
	} else if ret < 0 && n.left != nil {
		return n.left.find(key)
	} else if ret == 0 {
		return n
	} else {
		return nil
	}
}

func Tree_iterate(n *node, ch chan element) {
	if n == nil {
		return
	}
	Tree_iterate(n.left, ch)
	ch <- n.value
	Tree_iterate(n.right, ch)
}

func (n *node) leftleft() *node {
	temp := n.left
	n.left = temp.right
	temp.right = n
	n.get_height()
	temp.get_height()
	return temp
}

func (n *node) rightright() *node {
	temp := n.right
	n.right = temp.left
	temp.left = n
	n.get_height()
	temp.get_height()
	return temp
}

func (n *node) leftright() *node {
	temp := n.left
	n.left = temp.rightright()
	return n.leftleft()
}

func (n *node) rightleft() *node {
	temp := n.right
	n.left = temp.leftleft()
	return n.rightright()
}
