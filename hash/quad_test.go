package hash

import (
	util "benJames/util"
	"os"
)

func ExampleQuadInsert() {
	q := NewQuadHash(IntHash, 100)
	q.Insert(util.NewInteger(4))
	q.Display(os.Stdout)
	// Output: 4
}

func ExampleQuadInsertList() {
	q := NewQuadHash(IntHash, 100)
	q.InsertList(util.NewInteger(5), util.NewInteger(-4), util.NewInteger(32))
	q.Display(os.Stdout)
	// Output: -4
	// 5
	// 32
}

func ExampleQuadCollision() {
	q := NewQuadHash(IntHash, 100)
	q.InsertList(util.NewInteger(5), util.NewInteger(21))
	q.Display(os.Stdout)
	// Output: 5
	// 21
}

func ExampleQuadUpdate() {
	q := NewQuadHash(IntHash, 100)
	q.InsertList(util.NewInteger(5), util.NewInteger(-4), util.NewInteger(5))
	q.Display(os.Stdout)
	// Output: -4
	// 5
}

func ExampleQuadResize() {
	q := NewQuadHash(IntHash, 30)
	for i := 1; i <= 20; i++ {
		q.Insert(util.NewInteger(i))
	}
	q.Display(os.Stdout)
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
