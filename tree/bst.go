package tree

import u "benJames/util"

type bst struct {
	head *node
}

func (b *bst) init() *bst {
	b.head = nil
	return b
}

func NewBST() *bst {
	return new(bst).init()
}

func bst_insert(tree *node, value u.Element) *node {
	if tree == nil {
		return NewNode(value)
	}
	ret := value.Compare(tree.value)
	if ret == 0 {
		tree.value.Update()
	} else if ret > 0 {
		tree.right = bst_insert(tree.right, value)
	} else if ret < 0 {
		tree.left = bst_insert(tree.left, value)
	}
	return tree
}

func (b *bst) InsertList(values ...u.Element) {
	for _, v := range values {
		b.Insert(v)
	}
}

func (b *bst) Insert(value u.Element) {
	b.head = bst_insert(b.head, value)
}

func (b *bst) Find(key u.Element) u.Element {
	if key != nil && b.head != nil {
		n := b.head.find(key)
		if n != nil {
			return n.value
		}
	}
	return nil
}

func (b *bst) Display() {
	if b.head != nil {
		b.head.Display()
	}
}

func (b *bst) Iterator() <-chan u.Element {
	ch := make(chan u.Element)
	go func() {
		Tree_iterate(b.head, ch)
		close(ch)
	}()
	return ch
}
