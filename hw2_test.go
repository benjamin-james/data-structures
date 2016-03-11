package main

import (
	text "benJames/text"
	tree "benJames/tree"
	"os"
)

func ExampleCompute() {
	filename := "._temp"
	sample := "this is some sample text that I can use for this project"
	uni, bi, res := tree.NewAVL(), tree.NewAVL(), tree.NewAVL()
	assert(text.WriteString(filename, sample))
	Compute(filename, uni, bi, res)
	uni.Display(os.Stdout)
	bi.Display(os.Stdout)
	res.Display(os.Stdout)
	assert(os.Remove(filename))
	// Output: "I" 1
	// "can" 1
	// "for" 1
	// "is" 1
	// "project" 1
	// "sample" 1
	// "some" 1
	// "text" 1
	// "that" 1
	// "this" 2
	// "use" 1
	// "I can" 1
	// "can use" 1
	// "for this" 1
	// "is some" 1
	// "sample text" 1
	// "some sample" 1
	// "text that" 1
	// "that I" 1
	// "this is" 1
	// "this project" 1
	// "use for" 1
	// P(can|I) = 1.000000
	// P(use|can) = 1.000000
	// P(this|for) = 1.000000
	// P(some|is) = 1.000000
	// P(text|sample) = 1.000000
	// P(sample|some) = 1.000000
	// P(that|text) = 1.000000
	// P(I|that) = 1.000000
	// P(is|this) = 0.500000
	// P(project|this) = 0.500000
	// P(for|use) = 1.000000
}
