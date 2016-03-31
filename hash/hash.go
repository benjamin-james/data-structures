package hash

import (
	util "benJames/util"
)

type HashFunc func(util.Element, int) int

type KeySet struct {
	value util.Element
	next  *KeySet
}

func keyset_insert(ptr **KeySet, value util.Element) {
	tmp := *ptr
	var prev *KeySet
	for tmp != nil && tmp.value.Compare(value) <= 0 {
		prev, tmp = tmp, tmp.next
	}
	if prev == nil {
		*ptr = &KeySet{value, tmp}
	} else {
		prev.next = &KeySet{value, tmp}
	}
}
