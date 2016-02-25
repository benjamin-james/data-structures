package main

import (
	"fmt"
	"strings"
)

type prob struct {
	str1, str2 string
	num        float64
}

func NewProb(str1, str2 string, bi, uni int) *prob {
	return &prob{str1, str2, float64(bi) / float64(uni)}
}

func (a *prob) Compare(b element) int {
	if x, is_prob := b.(*prob); is_prob {
		a_str := fmt.Sprintf("%s %s", a.str1, a.str2)
		b_str := fmt.Sprintf("%s %s", x.str1, x.str2)
		return strings.Compare(a_str, b_str)
	} else {
		return 0
	}
}

func (p *prob) Update() {}

func (p *prob) String() string {
	return fmt.Sprintf("[\"%s %s\": %f]", p.str1, p.str2, p.num)
}

func ComputeProbabilities(unigrams, bigrams, result datastructure) {
	for x := range bigrams.Iterator() {
		if bi, is_bi := x.(*bigram); is_bi {
			y := unigrams.Find(NewUni(bi.str1, 1))
			if uni, is_uni := y.(*unigram); is_uni {
				result.Insert(NewProb(bi.str1, bi.str2, bi.count, uni.count))
			}
		}
	}
}
