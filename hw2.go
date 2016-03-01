package main

import (
	prob "benJames/problem"
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
	err := prob.ReadFile(filename, uni, bi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	prob.ComputeProbabilities(uni, bi, result)
	result.Display()
}
