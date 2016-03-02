package text

import (
	u "benJames/util"
	"fmt"
	"strings"
)

type Prob struct {
	str1, str2 string
	num        float64
}

func NewProb(str1, str2 string, bi, uni int) *Prob {
	return &Prob{str1, str2, float64(bi) / float64(uni)}
}

func (a *Prob) Compare(b u.Element) int {
	if x, is_Prob := b.(*Prob); is_Prob {
		a_str := fmt.Sprintf("%s %s", a.str1, a.str2)
		b_str := fmt.Sprintf("%s %s", x.str1, x.str2)
		return strings.Compare(a_str, b_str)
	} else {
		return 0
	}
}

func (p *Prob) Update() {}

func (p *Prob) String() string {
	return fmt.Sprintf("[\"%s %s\": %f]", p.str1, p.str2, p.num)
}

func ComputeProbabilities(unigrams, bigrams, result u.DataStructure) {
	for x := range bigrams.Iterator() {
		if bi, is_bi := x.(*Bigram); is_bi {
			y := unigrams.Find(NewUni(bi.str1, 1))
			if uni, is_uni := y.(*Unigram); is_uni {
				result.Insert(NewProb(bi.str1, bi.str2, bi.count, uni.count))
			}
		}
	}
}
