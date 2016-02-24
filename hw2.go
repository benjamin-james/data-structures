package main

import "fmt"

func main() {
	uni, bi := NewAVL(), NewAVL()
	err := ReadFile("/dev/stdin", uni, bi)
	if err == nil {
		fmt.Println("\nUnigrams")
		uni.Display()
		fmt.Println("\nBigrams")
		bi.Display()
	}
}
