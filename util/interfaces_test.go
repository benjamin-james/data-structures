package util

import (
	"fmt"
	"testing"
)

func TestIntegerCompare(t *testing.T) {
	a, b := NewInteger(2), NewInteger(2)
	if a.Compare(b) != 0 {
		t.Fatal("2 =/= 2")
	}
}

func ExampleIntegerString() {
	a := NewInteger(2)
	fmt.Println(a.String())
	// Output: 2
}
