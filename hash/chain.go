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
	ks    *KeySet
	hash  HashFunc
}

func NewChainHash(hash HashFunc, size int) *ChainHash {
	return &ChainHash{make([]*chain_elt, size), nil, hash}
}

func list_insert(list *chain_elt, e util.Element, ks **KeySet) *chain_elt {
	if list == nil {
		return &chain_elt{e, nil}
	} else if list.value.Compare(e) == 0 {
		list.value.Update()
		return list
	} else {
		return &chain_elt{e, list}
	}
}

func (c *ChainHash) Insert(e util.Element) {
	pos := c.hash(e, len(c.table))
	if c.table[pos] == nil {
		c.table[pos] = &chain_elt{e, nil}
		keyset_insert(&c.ks, e)
	} else if c.table[pos].value.Compare(e) == 0 {
		c.table[pos].value.Update()
	} else {
		c.table[pos] = &chain_elt{e, c.table[pos]}
		keyset_insert(&c.ks, e)
	}
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
