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
	uni, bi := NewAVL(), NewAVL()
	err := ReadFile(filename, uni, bi)
	if err == nil {
		fmt.Println("\nUnigrams")
		uni.Display()
		fmt.Println("\nBigrams")
		bi.Display()
	}
}
