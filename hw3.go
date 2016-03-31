package main

import (
	hash "benJames/hash"
	text "benJames/text"
	util "benJames/util"
	"fmt"
)

func strhash(e util.Element, size int) int {
	str := e.String()
	sum := size
	for _, char := range str {
		sum = (int(char) + 37*sum) % size
	}
	return sum
}

func main() {
	filename := "/dev/stdin"
	uni_file, bi_file, cp_file, time_file := "uni", "bi", "cp", "time"
	uni, bi, result := hash.NewChainHash(strhash, 1231231), hash.NewChainHash(strhash, 1231231), hash.NewChainHash(strhash, 1231231)
	ins_time, read_time, total_time, wc := Compute(filename, uni, bi, result)
	str := fmt.Sprintf("filename:\t\t%s\nword count:\t\t%d\nInsert time:\t\t%s\nRead time:\t\t%s\nTotal time\t\t%s\n", filename, wc, ins_time.String(), read_time.String(), total_time.String())
	text.DumpToFile(uni_file, uni)
	text.DumpToFile(bi_file, bi)
	text.DumpToFile(cp_file, result)
	text.WriteString(time_file, str)

}
