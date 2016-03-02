package tree

import (
	u "benJames/util"
)

func ExampleBSTInsert() {
	b := NewBST()
	b.Insert(u.NewInteger(4))
	b.Display()
	// Output: 4
}

func ExampleBSTInsertList() {
	b := NewBST()
	b.InsertList(u.NewInteger(5), u.NewInteger(-4), u.NewInteger(32))
	b.Display()
	// Output: -4
	// 5
	// 32
}

func ExampleBSTCollision() {
	b := NewBST()
	b.InsertList(u.NewInteger(5), u.NewInteger(-4), u.NewInteger(5))
	b.Display()
	// Output: -4
	// 5
}
