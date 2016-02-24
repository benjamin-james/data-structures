package main

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

func bst_insert(tree *node, value element) *node {
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

func (b *bst) Insert(value element) {
	b.head = bst_insert(b.head, value)
}

func (b *bst) Find(key element) element {
	if key != nil && b.head != nil {
		n := b.head.find(key)
		if n != nil {
			return n.value
		}
	}
	return nil
}
