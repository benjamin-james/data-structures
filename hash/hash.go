package hash

import (
	util "benJames/util"
)

type HashFunc func(util.Element, int) int

type HashElt struct {
	value  util.Element
	exists bool
}

type QuadHash struct {
	table    []HashElt
	capacity int
	hash     HashFunc
}

func NewQuadHash(hash HashFunc, size int) *QuadHash {
	return &QuadHash{make([]HashElt, size), 0, hash}
}

func (h *QuadHash) f(e util.Element, i int) {
	i %= len(h.table)
	if h.table[i].exists {
		if h.table[i].value.Compare(e) == 0 {
			h.table[i].value.Update()
		} else {
			h.f(e, i*i)
		}

	} else {
		h.table[i].value = e
		h.table[i].exists = true
		h.capacity++
	}
}

func (h *QuadHash) Insert(e util.Element) {
	pos := h.hash(e, len(h.table))
	h.f(e, pos)
	if h.capacity >= len(h.table)/2 {
		//resize, rehash
	}
}
