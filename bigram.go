package main

import (
	"fmt"
	"strings"
)

type bigram struct {
	str1, str2 string
	count      int
}

func NewBi(str1, str2 string, count int) bigram {
	return bigram{str1, str2, count}
}

func Bi_cmp(a, b interface{}) int {
	a_str, b_str := "", ""
	if x, err_x := a.(bigram); err_x {
		a_str = fmt.Sprintf("%s %s", x.str1, x.str2)
	}
	if y, err_y := b.(bigram); err_y {
		b_str = fmt.Sprintf("%s %s", y.str1, y.str2)
	}
	return strings.Compare(a_str, b_str)
}

func Bi_eq(v interface{}) interface{} {
	if value, err_v := v.(bigram); err_v {
		value.count++
		return value
	}
	return nil
}
