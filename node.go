package main

import "fmt"

type cmpfunc func(a interface{}, b interface{}) int
type eq_hook func(a interface{}) interface{}

type node struct {
	value       interface{} // Go's way out of templates
	left, right *node
	height      int
}

func (n *node) init(value interface{}) *node {
	n.left = nil
	n.right = nil
	n.height = 1
	n.value = value
	return n
}

func NewNode(value interface{}) *node {
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

func (n *node) find(key interface{}, cmp cmpfunc) *node {
	ret := cmp(key, n.value)
	if ret > 0 && n.right != nil {
		return n.right.find(key, cmp)
	} else if ret < 0 && n.left != nil {
		return n.left.find(key, cmp)
	} else if ret == 0 {
		return n
	} else {
		return nil
	}
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
