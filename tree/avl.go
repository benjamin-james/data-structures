package tree

import (
	u "benJames/util"
	"io"
)

type avl struct {
	head *node
}

func (a *avl) init() *avl {
	a.head = nil
	return a
}

func (a *avl) Display(w io.Writer) {
	if a.head != nil {
		a.head.Display(w)
	}
}

func NewAVL() *avl {
	return new(avl).init()
}

func balance(tree *node) *node {
	bf := tree.diff()
	if bf > 1 && tree.left != nil {
		if tree.left.diff() > 0 {
			tree = tree.leftleft()
		} else {
			tree = tree.leftright()
		}
	} else if bf < -1 && tree.right != nil {
		if tree.right.diff() > 0 {
			tree = tree.rightleft()
		} else {
			tree = tree.rightright()
		}
	}
	return tree
}

func insert(tree *node, value u.Element) *node {
	if tree == nil {
		return NewNode(value)
	}
	compare := value.Compare(tree.value)
	if compare == 0 {
		tree.value.Update()
	} else if compare > 0 {
		tree.right = insert(tree.right, value)
		tree.get_height()
		tree = balance(tree)
	} else if compare < 0 {
		tree.left = insert(tree.left, value)
		tree.get_height()
		tree = balance(tree)
	}
	return tree
}

func (a *avl) Insert(value u.Element) {
	a.head = insert(a.head, value)
}

func (a *avl) InsertList(values ...u.Element) {
	for _, v := range values {
		a.Insert(v)
	}
}

func (a *avl) Find(key u.Element) u.Element {
	if key != nil && a.head != nil {
		n := a.head.find(key)
		if n != nil {
			return n.value
		}
	}
	return nil
}

func (a *avl) Iterator() <-chan u.Element {
	ch := make(chan u.Element)
	go func() {
		Tree_iterate(a.head, ch)
		close(ch)
	}()
	return ch
}
