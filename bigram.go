package main

import (
	"fmt"
	"strings"
)

type bigram struct {
	str1, str2 string
	count      int
}

func NewBi(str1, str2 string, count int) *bigram {
	return &bigram{str1, str2, count}
}

func (a *bigram) Compare(b element) int {
	if x, is_bi := b.(*bigram); is_bi {
		a_str := fmt.Sprintf("%s %s", a.str1, a.str2)
		b_str := fmt.Sprintf("%s %s", x.str1, x.str2)
		return strings.Compare(a_str, b_str)
	} else {
		return 0
	}
}

func (b *bigram) Update() {
	b.count++
}

func (b *bigram) String() string {
	return fmt.Sprintf("[\"%s %s\": %d]", b.str1, b.str2, b.count)
}
