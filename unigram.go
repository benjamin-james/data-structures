package main

import (
	"strings"
)

type unigram struct {
	str   string
	count int
}

func NewUni(str string, count int) unigram {
	return unigram{str, count}
}

func Uni_cmp(a, b interface{}) int {
	a_str, b_str := "", ""
	if x, err_x := a.(unigram); err_x {
		a_str = x.str
	}
	if y, err_y := b.(unigram); err_y {
		b_str = y.str
	}
	return strings.Compare(a_str, b_str)
}

func Uni_eq(v interface{}) interface{} {
	if value, err_v := v.(unigram); err_v {
		value.count++
		return value
	}
	return nil
}
