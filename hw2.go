package main

import (
	text "benJames/text"
	tree "benJames/tree"
	"fmt"
	"os"
)

func main() {
	filename := "/dev/stdin"
	uni_file, bi_file, cp_file, time_file := "uni", "bi", "cp", "time"
	if len(os.Args) > 1 {
		filename = os.Args[1]
		uni_file = fmt.Sprintf("%s.%s", filename, uni_file)
		bi_file = fmt.Sprintf("%s.%s", filename, bi_file)
		cp_file = fmt.Sprintf("%s.%s", filename, cp_file)
		time_file = fmt.Sprintf("%s.%s", filename, time_file)
	}
	uni, bi, result := tree.NewAVL(), tree.NewAVL(), tree.NewAVL()
	err := text.ReadFile(filename, uni, bi)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}
	text.ComputeProbabilities(uni, bi, result)
}
