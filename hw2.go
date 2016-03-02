package main

import (
	text "benJames/text"
	tree "benJames/tree"
	"fmt"
	"os"
)

func main() {

	filename := "/dev/stdin"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	uni, bi, result := tree.NewAVL(), tree.NewAVL(), tree.NewAVL()
	err := text.ReadFile(filename, uni, bi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	text.ComputeProbabilities(uni, bi, result)
	result.Display()
}
