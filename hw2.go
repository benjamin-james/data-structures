package main

import (
	"fmt"
	"os"
)

func main() {

	filename := "/dev/stdin"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	uni, bi, result := NewBST(), NewBST(), NewBST()
	err := ReadFile(filename, uni, bi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	ComputeProbabilities(uni, bi, result)
	result.Display()
}
