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
	hash  HashFunc
}

func NewChainHash(hash HashFunc, size int) *ChainHash {
	return &ChainHash{make([]*chain_elt, size), hash}
}

func list_insert(list *chain_elt, e util.Element) *chain_elt {
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
	c.table[pos] = list_insert(c.table[pos], e)
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
	for _, e := range c.table {
		for elt := e; elt != nil; elt = elt.next {
			fmt.Fprintln(w, elt.value)
		}
	}
}

func (c *ChainHash) Iterator() <-chan util.Element {
	ch := make(chan util.Element)
	go func() {
		for _, e := range c.table {
			for ; e != nil; e = e.next {
				ch <- e.value
			}
		}
		close(ch)
	}()
	return ch
}
