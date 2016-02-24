package main

import (
	"fmt"
)

func main() {
	avl := NewAVL()
	avl.Insert(NewUni("foo", 1))
	avl.Insert(NewUni("bar", 1))
	avl.Insert(NewUni("baz", 1))
	avl.Insert(NewUni("foo", 1))
	avl.Display()
	fmt.Println("The value of foo is", avl.Find(NewUni("foo", 1)))
}
