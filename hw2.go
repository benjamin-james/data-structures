package main

import (
	text "benJames/text"
	tree "benJames/tree"
	util "benJames/util"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	filename := "/dev/stdin"
	uni_file, bi_file, cp_file, time_file := "uni", "bi", "cp", "time"
	str := ""
	use_bst := true
	use_avl := true
	has_input := false
	for argc, argv := range os.Args {
		if argc == 0 {
			continue
		} else if strings.Compare("-h", argv) == 0 || strings.Compare("--help", argv) == 0 {
			fmt.Printf("Usage: %s [--all|--bst|--no-bst|--avl|--no-avl] filename\n", os.Args[0])
			os.Exit(0)
		} else if strings.Compare("--all", argv) == 0 || strings.Compare("-a", argv) == 0 {
			use_avl = true
			use_bst = true
		} else if strings.Compare("--bst", argv) == 0 {
			use_bst = true
		} else if strings.Compare("--no-bst", argv) == 0 {
			use_bst = false
		} else if strings.Compare("--avl", argv) == 0 {
			use_avl = true
		} else if strings.Compare("--no-avl", argv) == 0 {
			use_avl = false
		} else if _, err := os.Stat(argv); err == nil && !has_input {
			filename = argv
			has_input = true
			uni_file = fmt.Sprintf("%s.%s", filename, uni_file)
			bi_file = fmt.Sprintf("%s.%s", filename, bi_file)
			cp_file = fmt.Sprintf("%s.%s", filename, cp_file)
			time_file = fmt.Sprintf("%s.%s", filename, time_file)
		} else if !has_input {
			fmt.Fprintf(os.Stderr, "The file \"%s\" does not exist.\n", argv)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stderr, "Bad argument \"%s\"\n", argv)
			os.Exit(1)
		}
	}
	if use_avl {
		uni, bi, result := tree.NewAVL(), tree.NewAVL(), tree.NewAVL()
		dur, wc := Compute(filename, uni, bi, result)
		if use_bst {
			uni_bst, bi_bst, result_bst := tree.NewBST(), tree.NewBST(), tree.NewBST()
			dur_bst, _ := Compute(filename, uni_bst, bi_bst, result_bst)
			str = fmt.Sprintf("filename:\t%s\nword count:\t%d\nAVL time:\t%s\nBST time:\t%s\n", filename, wc, dur.String(), dur_bst.String())
		} else {
			str = fmt.Sprintf("filename:\t%s\nword count:\t%d\nAVL time:\t%s\n", filename, wc, dur.String())
		}
		text.DumpToFile(uni_file, uni)
		text.DumpToFile(bi_file, bi)
		text.DumpToFile(cp_file, result)
		text.WriteString(time_file, str)
	} else if use_bst {
		uni, bi, result := tree.NewBST(), tree.NewBST(), tree.NewBST()
		dur, wc := Compute(filename, uni, bi, result)
		str = fmt.Sprintf("filename:\t%s\nword count:\t%d\nBST time:\t%s\n", filename, wc, dur.String())
		text.DumpToFile(uni_file, uni)
		text.DumpToFile(bi_file, bi)
		text.DumpToFile(cp_file, result)
		text.WriteString(time_file, str)
	}
}

func Compute(filename string, uni, bi, result util.DataStructure) (time.Duration, int) {
	now := time.Now()
	wc, err := text.ReadFile(filename, uni, bi)
	assert(err)
	text.ComputeProbabilities(uni, bi, result)
	return time.Since(now), wc
}

// similar to assert(3)
func assert(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
