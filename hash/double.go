package hash

import (
	util "benJames/util"
	"fmt"
	"io"
)

type DoubleHash struct {
	table  []*util.Element
	size   int
	ks     *KeySet
	h1, h2 HashFunc
}

func NewDoubleHash(h1, h2 HashFunc, size int) *DoubleHash {
	return &DoubleHash{make([]*util.Element, size), 0, nil, h1, h2}
}

func (d *DoubleHash) Display(w io.Writer) {
	for x := range d.Iterator() {
		fmt.Fprintln(w, x)
	}
}

func (d *DoubleHash) Iterator() <-chan util.Element {
	ch := make(chan util.Element)
	go func() {
		for ks := d.ks; ks != nil; ks = ks.next {
			e := d.Get(ks.value)
			if e != nil {
				ch <- *e
			}
		}
		close(ch)
	}()
	return ch
}

func (d *DoubleHash) Get(e util.Element) *util.Element {
	hash1, hash2 := d.h1(e, len(d.table)), d.h2(e, len(d.table)-1)+1
	for i := 0; i < len(d.table); i++ {
		index := (hash1 + i*hash2) % len(d.table)
		if d.table[index] != nil && e.Compare(*d.table[index]) == 0 {
			return d.table[index]
		}
	}
	return nil
}

func (d *DoubleHash) Insert(e util.Element) {
	hash1, hash2 := d.h1(e, len(d.table)), d.h2(e, len(d.table)-1)+1
	for i := 0; i < len(d.table); i++ {
		index := (hash1 + i*hash2) % len(d.table)
		if d.table[index] != nil && e.Compare(*d.table[index]) == 0 {
			(*d.table[index]).Update()
			break
		} else if d.table[index] == nil {
			keyset_insert(&d.ks, e)
			d.table[index] = &e
			d.size++
			if d.size >= len(d.table)/2 {
				d.ResizeAndRehash()
			}
			break
		}
	}
}

func (d *DoubleHash) ResizeAndRehash() {
	table := make([]*util.Element, len(d.table)*2+1)
	for ks := d.ks; ks != nil; ks = ks.next {
		hash1, hash2 := d.h1(ks.value, len(table)), d.h2(ks.value, len(table)-1)+1
		for i := 0; i < len(table); i++ {
			index := (hash1 + i*hash2) % len(table)
			if table[index] == nil {
				table[index] = d.Get(ks.value)
				break
			}
		}
	}
	d.table = table
}

func (d *DoubleHash) InsertList(vars ...util.Element) {
	for _, v := range vars {
		d.Insert(v)
	}
}

func (d *DoubleHash) Find(e util.Element) util.Element {
	return *d.Get(e)
}
