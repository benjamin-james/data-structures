package text

import (
	tree "benJames/tree"
	"os"
)

func ExampleComputeProbabilities() {
	uni, bi, res := tree.NewAVL(), tree.NewAVL(), tree.NewAVL()
	uni.InsertList(NewUni("foo", 1), NewUni("bar", 1), NewUni("foo", 1), NewUni("baz", 1))
	bi.InsertList(NewBi("foo", "bar", 1), NewBi("bar", "foo", 1), NewBi("foo", "baz", 1))
	ComputeProbabilities(uni, bi, res)
	res.Display(os.Stdout)
	// Output: ["bar foo": 1.000000]
	// ["foo bar": 0.500000]
	// ["foo baz": 0.500000]
}
