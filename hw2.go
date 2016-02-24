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
	uni, bi := NewBST(), NewBST()
	err := ReadFile(filename, uni, bi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	fmt.Println("\nUnigrams")
	uni.Display()
	fmt.Println("\nBigrams")
	bi.Display()
}
