package hash

import (
	util "benJames/util"
	"fmt"
	"io"
)

type QuadHash struct {
	table []*util.Element
	size  int
	ks    *KeySet
	hash  HashFunc
}

func NewQuadHash(hash HashFunc, size int) *QuadHash {
	return &QuadHash{make([]*util.Element, size), 0, nil, hash}
}

func (h *QuadHash) Display(w io.Writer) {
	for ks := h.ks; ks != nil; ks = ks.next {
		e := h.Get(ks.value)
		if e == nil {
			continue
		}
		fmt.Fprintln(w, *e)
	}
}

func (h *QuadHash) InsertList(values ...util.Element) {
	for _, v := range values {
		h.Insert(v)
	}
}

func (h *QuadHash) Find(e util.Element) util.Element {
	v := h.Get(e)
	if v == nil {
		return e
	} else {
		return *v
	}
}

func (h *QuadHash) Iterator() <-chan util.Element {
	ch := make(chan util.Element)
	go func() {
		for ks := h.ks; ks != nil; ks = ks.next {
			e := h.Get(ks.value)
			if e != nil {
				ch <- *e
			}
		}
		close(ch)
	}()
	return ch
}

func (h *QuadHash) Insert(e util.Element) {
	hash := h.hash(e, len(h.table))
	for i := 0; i < len(h.table); i++ {
		index := (hash + i*i) % len(h.table)
		if h.table[index] != nil && e.Compare(*h.table[index]) == 0 {
			(*h.table[index]).Update()
			break
		} else if h.table[index] == nil {
			keyset_insert(&h.ks, e)
			h.table[index] = &e
			h.size++
			if h.size >= len(h.table)/2 {
				h.ResizeAndRehash()
			}
			break
		}
	}

}

func (h *QuadHash) ResizeAndRehash() {
	table := make([]*util.Element, len(h.table)*2+1)
	for ks := h.ks; ks != nil; ks = ks.next {
		hash := h.hash(ks.value, len(table))
		for i := 0; i < len(table); i++ {
			index := (hash + i*i) % len(table)
			if table[index] == nil {
				table[index] = h.Get(ks.value)
				break
			}
		}
	}
	h.table = table
}

func (h *QuadHash) Get(e util.Element) *util.Element {
	hash := h.hash(e, len(h.table))
	for i := 0; i < len(h.table); i++ {
		index := (hash + i*i) % len(h.table)
		if h.table[index] != nil && e.Compare(*h.table[index]) == 0 {
			return h.table[index]
		}
	}
	return nil
}
