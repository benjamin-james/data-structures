package main

import (
	"fmt"
	"strings"
)

type data struct {
	str   string
	count int
}

func mycmp(a, b interface{}) int {
	a_str := ""
	b_str := ""
	if x, err_x := a.(data); err_x {
		a_str = x.str
	}
	if y, err_y := b.(data); err_y {
		b_str = y.str
	}
	return strings.Compare(a_str, b_str)
}

func my_eq(v interface{}) interface{} {

	if value, err_v := v.(data); err_v {
		value.count++
		return value
	} else {
		return v
	}
}
func main() {
	avl := NewAVL(mycmp, my_eq)
	avl.Insert(data{"foo", 1})
	avl.Insert(data{"bar", 1})
	avl.Insert(data{"baz", 1})
	avl.Insert(data{"foo", 1})
	avl.Display()
	fmt.Println("The value of foo is", avl.Find(data{"foo", 0}))
}
