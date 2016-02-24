package main

import (
	"fmt"
	"strings"
)

type unigram struct {
	str   string
	count int
}

func NewUni(str string, count int) *unigram {
	return &unigram{str, count}
}

func (a *unigram) Compare(b element) int {
	if x, is_uni := b.(*unigram); is_uni {
		return strings.Compare(a.str, x.str)
	} else {
		return 0
	}
}

func (u *unigram) Update() {
	u.count++
}

func (u *unigram) String() string {
	return fmt.Sprintf("[\"%s\": %d]", u.str, u.count)
}
