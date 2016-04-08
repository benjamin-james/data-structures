package main

import (
	hash "benJames/hash"
	text "benJames/text"
	util "benJames/util"
	"fmt"
	"os"
)

func strhash2(e util.Element, size int) int {
	str := e.String()
	var sum uint64 = 3219
	for _, char := range str {
		sum = sum*37 + uint64(char)
	}
	return int(sum % uint64(size))
}
func strhash(e util.Element, size int) int {
	str := e.String()
	var sum uint64 = 152501
	for _, char := range str {
		sum = (sum << 5) + sum + uint64(char) /* sum = 33 * sum + char */
	}
	return int(sum % uint64(size))
}

func main() {
	filename := "/dev/stdin"
	has_input := false
	uni_file, bi_file, cp_file, time_file := "uni", "bi", "cp", "time"
	hashsize := 324123
	for argc, argv := range os.Args {
		if argc == 0 {
			continue
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
	chain_uni, chain_bi, chain_result := hash.NewChainHash(strhash, hashsize), hash.NewChainHash(strhash, hashsize), hash.NewChainHash(strhash, hashsize)
	chain_ins_time, chain_read_time, chain_total_time, wc := Compute(filename, chain_uni, chain_bi, chain_result)
	str := fmt.Sprintf("filename:\t\t%s\nword count:\t\t%d\nChain insert time:\t%s\nChain read time:\t%s\nChain total time:\t%s\n", filename, wc, chain_ins_time.String(), chain_read_time.String(), chain_total_time.String())
	text.DumpToFile(uni_file, chain_uni)
	text.DumpToFile(bi_file, chain_bi)
	text.DumpToFile(cp_file, chain_result)
	quad_uni, quad_bi, quad_result := hash.NewQuadHash(strhash, hashsize), hash.NewQuadHash(strhash, hashsize), hash.NewChainHash(strhash, hashsize)
	quad_ins_time, quad_read_time, quad_total_time, _ := Compute(filename, quad_uni, quad_bi, quad_result)
	str = fmt.Sprintf("%sQuad insert time:\t%s\nQuad read time:\t\t%s\nQuad total time:\t%s\n", str, quad_ins_time.String(), quad_read_time.String(), quad_total_time.String())

	dbl_uni, dbl_bi, dbl_result := hash.NewDoubleHash(strhash, strhash2, hashsize), hash.NewDoubleHash(strhash, strhash2, hashsize), hash.NewDoubleHash(strhash, strhash2, hashsize)
	dbl_ins_time, dbl_read_time, dbl_total_time, _ := Compute(filename, dbl_uni, dbl_bi, dbl_result)
	str = fmt.Sprintf("%sDouble insert time:\t%s\nDouble read time:\t%s\nDouble total time:\t%s\n", str, dbl_ins_time.String(), dbl_read_time.String(), dbl_total_time.String())
	text.WriteString(time_file, str)
}
