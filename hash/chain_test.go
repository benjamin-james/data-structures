package hash

import (
	util "benJames/util"
	"os"
)

func IntHash(e util.Element, size int) int {
	if i, is_int := e.(*util.Integer); is_int {
		return i.Num % size
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
