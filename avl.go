package main

type avl struct {
	head *node
	cmp  cmpfunc
	eq   eq_hook
}

func (a *avl) init(cmp cmpfunc, eq eq_hook) *avl {
	a.head = nil
	a.cmp = cmp
	a.eq = eq
	return a
}

func (a *avl) Display() {
	if a.head != nil {
		a.head.Display()
	}
}

func NewAVL(cmp cmpfunc, eq eq_hook) *avl {
	return new(avl).init(cmp, eq)
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

func insert(tree *node, value interface{}, cmp cmpfunc, eq eq_hook) *node {
	if tree == nil {
		return NewNode(value)
	}
	compare := cmp(value, tree.value)
	if compare == 0 {
		tree.value = eq(tree.value)
	} else if compare > 0 {
		tree.right = insert(tree.right, value, cmp, eq)
		tree.get_height()
		tree = balance(tree)
	} else if compare < 0 {
		tree.left = insert(tree.left, value, cmp, eq)
		tree.get_height()
		tree = balance(tree)
	}
	return tree
}

func (a *avl) Insert(value interface{}) {
	a.head = insert(a.head, value, a.cmp, a.eq)
}

func (a *avl) Find(key interface{}) interface{} {
	if key != nil && a.head != nil {
		n := a.head.find(key, a.cmp)
		if n != nil {
			return n.value
		}
	}
	return nil
}