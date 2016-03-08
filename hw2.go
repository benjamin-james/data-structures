package main

import (
	text "benJames/text"
	tree "benJames/tree"
	util "benJames/util"
	"fmt"
	"os"
	"time"
)

func Compute(filename string, uni, bi, result util.DataStructure) (time.Duration, int) {
	now := time.Now()
	wc, err := text.ReadFile(filename, uni, bi)
	assert(err)
	text.ComputeProbabilities(uni, bi, result)
	return time.Since(now), wc
}

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
	uni_avl, bi_avl, result_avl := tree.NewAVL(), tree.NewAVL(), tree.NewAVL()
	uni_bst, bi_bst, result_bst := tree.NewBST(), tree.NewBST(), tree.NewBST()
	dur_avl, wc := Compute(filename, uni_avl, bi_avl, result_avl)
	dur_bst, _ := Compute(filename, uni_bst, bi_bst, result_bst)
	assert(text.DumpToFile(uni_file, uni_avl))
	assert(text.DumpToFile(bi_file, bi_avl))
	assert(text.DumpToFile(cp_file, result_avl))
	s := fmt.Sprintf("filename:\t%s\nword count:\t%d\nAVL time:\t%s\nBST time:\t%s", filename, wc, dur_avl.String(), dur_bst.String())
	text.WriteString(time_file, s)
}

// similar to assert(3)
func assert(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
