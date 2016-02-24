package main

type bst struct {
	head *node
	cmp  cmpfunc
}

func (b *bst) init(cmp cmpfunc) *bst {
	b.head = nil
	b.cmp = cmp
	return b
}

func NewBST(cmp cmpfunc) *bst {
	return new(bst).init(cmp)
}

func bst_insert(tree *node, value interface{}, cmp cmpfunc) *node {
	if tree == nil {
		tree = NewNode(value)
	} else if cmp(value, tree.value) > 0 {
		tree.right = bst_insert(tree.right, value, cmp)
	} else if cmp(value, tree.value) < 0 {
		tree.left = bst_insert(tree.left, value, cmp)
	}
	return tree
}

func (b *bst) Insert(value interface{}) {
	b.head = bst_insert(b.head, value, b.cmp)
}

func (b *bst) Find(key interface{}) interface{} {
	if key != nil && b.head != nil {
		n := b.head.find(key, b.cmp)
		if n != nil {
			return n.value
		}
	}
	return nil
}
