package hash

import (
	util "benJames/util"
	"fmt"
	"io"
)

type chain_elt struct {
	value util.Element
	next  *chain_elt
}

type ChainHash struct {
	table []*chain_elt
	size  int
	ks    *KeySet
	hash  HashFunc
}

func NewChainHash(hash HashFunc, size int) *ChainHash {
	return &ChainHash{make([]*chain_elt, size), 0, nil, hash}
}

func (c *ChainHash) Insert(e util.Element) {
	pos := c.hash(e, len(c.table))
	if c.table[pos] == nil {
		c.table[pos] = &chain_elt{e, nil}
		c.size++
		keyset_insert(&c.ks, e)
	} else if c.table[pos].value.Compare(e) == 0 {
		c.table[pos].value.Update()
	} else {
		c.table[pos] = &chain_elt{e, c.table[pos]}
		c.size++
		keyset_insert(&c.ks, e)
	}
	if c.size >= len(c.table) {
		c.ResizeAndRehash()
	}
}

func (c *ChainHash) ResizeAndRehash() {
	table := make([]*chain_elt, len(c.table)*2+1)
	for ks := c.ks; ks != nil; ks = ks.next {
		hash := c.hash(ks.value, len(table))
		table[hash] = &chain_elt{ks.value, table[hash]}
	}
	c.table = table
}

func (c *ChainHash) InsertList(values ...util.Element) {
	for _, v := range values {
		c.Insert(v)
	}
}

func (c *ChainHash) Find(e util.Element) util.Element {
	pos := c.hash(e, len(c.table))
	for elt := c.table[pos]; elt != nil; elt = elt.next {
		if elt.value.Compare(e) == 0 {
			return elt.value
		}
	}
	return nil
}

func (c *ChainHash) Display(w io.Writer) {
	for x := range c.Iterator() {
		fmt.Fprintln(w, x)
	}
}

func (c *ChainHash) Iterator() <-chan util.Element {
	ch := make(chan util.Element)
	go func() {
		for ks := c.ks; ks != nil; ks = ks.next {
			ch <- c.Find(ks.value)
		}
		close(ch)
	}()
	return ch
}
