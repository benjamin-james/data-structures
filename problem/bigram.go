package problem

import (
	u "benJames/util"
	"fmt"
	"strings"
)

type Bigram struct {
	str1, str2 string
	count      int
}

func NewBi(str1, str2 string, count int) *Bigram {
	return &Bigram{str1, str2, count}
}

func (a *Bigram) Compare(b u.Element) int {
	if x, is_bi := b.(*Bigram); is_bi {
		a_str := fmt.Sprintf("%s %s", a.str1, a.str2)
		b_str := fmt.Sprintf("%s %s", x.str1, x.str2)
		return strings.Compare(a_str, b_str)
	} else {
		return 0
	}
}

func (b *Bigram) Update() {
	b.count++
}

func (b *Bigram) String() string {
	return fmt.Sprintf("[\"%s %s\": %d]", b.str1, b.str2, b.count)
}
