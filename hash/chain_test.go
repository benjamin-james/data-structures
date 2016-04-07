package hash

import (
	util "benJames/util"
	"os"
)

func IntHash(e util.Element, size int) int {
	if i, is_int := e.(*util.Integer); is_int {
		return ((i.Num & 0xF) + size) % size
	} else {
		return 0
	}
}

func ExampleChainInsert() {
	c := NewChainHash(IntHash, 100)
	c.Insert(util.NewInteger(4))
	c.Display(os.Stdout)
	// Output: 4
}

func ExampleChainInsertList() {
	c := NewChainHash(IntHash, 100)
	c.InsertList(util.NewInteger(5), util.NewInteger(-4), util.NewInteger(32))
	c.Display(os.Stdout)
	// Output: -4
	// 5
	// 32
}

func ExampleChainCollision() {
	c := NewChainHash(IntHash, 100)
	c.InsertList(util.NewInteger(5), util.NewInteger(21))
	c.Display(os.Stdout)
	// Output: 5
	// 21
}

func ExampleChainUpdate() {
	c := NewChainHash(IntHash, 100)
	c.InsertList(util.NewInteger(5), util.NewInteger(-4), util.NewInteger(5))
	c.Display(os.Stdout)
	// Output: -4
	// 5
}

func ExampleChainResizeAndRehash() {
	c := NewChainHash(IntHash, 5)
	for i := 1; i <= 20; i++ {
		c.Insert(util.NewInteger(i))
	}
	c.Display(os.Stdout)
	// Output: 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
	// 11
	// 12
	// 13
	// 14
	// 15
	// 16
	// 17
	// 18
	// 19
	// 20
}
