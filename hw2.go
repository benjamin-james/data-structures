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
	assert(text.ReadFile(filename, uni, bi))
	assert(text.DumpToFile(uni_file, uni))
	assert(text.DumpToFile(bi_file, bi))
	text.ComputeProbabilities(uni, bi, result)
	assert(text.DumpToFile(cp_file, result))
}

// similar to assert(3)
func assert(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
