package text

import (
	u "benJames/util"
	"fmt"
	"strings"
)

type Unigram struct {
	str   string
	count int
}

func NewUni(str string, count int) *Unigram {
	return &Unigram{str, count}
}

func (a *Unigram) Compare(b u.Element) int {
	if x, is_uni := b.(*Unigram); is_uni {
		return strings.Compare(a.str, x.str)
	} else {
		return 0
	}
}

func (u *Unigram) Update() {
	u.count++
}

func (u *Unigram) String() string {
	return fmt.Sprintf("[\"%s\": %d]", u.str, u.count)
}
