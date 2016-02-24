package main

import "strings"

type data struct {
	str   string
	count int
}

func mycmp(a, b interface{}) int {
	x := a.(data)
	y := b.(data)
	return strings.Compare(x.str, y.str)
}
func main() {
	avl := NewAVL(mycmp)
	avl.Insert(data{"foo", 1})
	avl.Insert(data{"bar", 1})
	avl.Insert(data{"baz", 1})
	avl.Display()
}
