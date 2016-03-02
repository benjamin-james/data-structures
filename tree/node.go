package tree

import (
	u "benJames/util"
	"fmt"
	"io"
)

type node struct {
	value       u.Element
	left, right *node
	height      int
}

func (n *node) init(value u.Element) *node {
	n.left = nil
	n.right = nil
	n.height = 1
	n.value = value
	return n
}

func NewNode(value u.Element) *node {
	return new(node).init(value)
}

func (n *node) Display(w io.Writer) {
	if n.left != nil {
		n.left.Display(w)
	}
	fmt.Fprintln(w, n.value)
	if n.right != nil {
		n.right.Display(w)
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

func (n *node) find(key u.Element) *node {
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

func Tree_iterate(n *node, ch chan u.Element) {
	if n == nil {
		return
	}
	Tree_iterate(n.left, ch)
	ch <- n.value
	Tree_iterate(n.right, ch)
}

func (n *node) Copy() *node {
	return &node{n.value, n.left, n.right, n.height}
}

func (n *node) leftleft() *node {
	new_n := n.left.Copy()
	n.left = new_n.right
	new_n.right = n
	n.get_height()
	new_n.get_height()
	return new_n
}

func (n *node) rightright() *node {
	new_n := n.right.Copy()
	n.right = new_n.left
	new_n.left = n
	n.get_height()
	new_n.get_height()
	return new_n
}

func (n *node) leftright() *node {
	temp := n.left
	n.left = temp.rightright()
	return n.leftleft()
}

func (n *node) rightleft() *node {
	temp := n.right
	n.right = temp.leftleft()
	return n.rightright()
}
