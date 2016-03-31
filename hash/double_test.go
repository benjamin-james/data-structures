package hash

import (
	util "benJames/util"
	"os"
)

func IntHash2(e util.Element, size int) int {
	if i, is_int := e.(*util.Integer); is_int {
		return (((i.Num * 5) & 0xF) + size) % size
	} else {
		return 0
	}
}

func ExampleDoubleInsert() {
	d := NewDoubleHash(IntHash, IntHash2, 100)
	d.Insert(util.NewInteger(4))
	d.Display(os.Stdout)
	// Output: 4
}

func ExampleDoubleInsertList() {
	d := NewDoubleHash(IntHash, IntHash2, 100)
	d.InsertList(util.NewInteger(5), util.NewInteger(21))
	d.Display(os.Stdout)
	// Output: 5
	// 21
}

func ExampleDoubleUpdate() {
	d := NewDoubleHash(IntHash, IntHash2, 100)
	d.InsertList(util.NewInteger(5), util.NewInteger(-4), util.NewInteger(5))
	d.Display(os.Stdout)
	// Output: -4
	// 5
}

func ExampleDoubleCollision() {
	d := NewDoubleHash(IntHash, IntHash2, 100)
	d.InsertList(util.NewInteger(5), util.NewInteger(21))
	d.Display(os.Stdout)
	// Output: 5
	// 21
}

func ExampleDoubleResize() {
	d := NewDoubleHash(IntHash, IntHash2, 30)
	for i := 1; i <= 20; i++ {
		d.Insert(util.NewInteger(i))
	}
	d.Display(os.Stdout)
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
