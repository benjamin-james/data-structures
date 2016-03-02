package util

import (
	"fmt"
)

type Element interface {
	Compare(Element) int
	Update()
	String() string
}

type DataStructure interface {
	Display()
	Insert(Element)
	InsertList(Element)
	Find(Element) Element
	Iterator() <-chan Element
}

/*
 * Sample struct to play with
 */
type Integer struct {
	num int
}

func NewInteger(n int) *Integer {
	return &Integer{n}
}

func (a *Integer) Compare(e Element) int {
	if b, is_int := e.(*Integer); is_int {
		return a.num - b.num
	} else {
		return 0
	}
}

func (i *Integer) Update() {}

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.num)
}