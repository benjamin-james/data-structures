package tree

import (
	u "benJames/util"
	"testing"
)

func ExampleAVLInsert() {
	a := NewAVL()
	a.Insert(u.NewInteger(4))
	a.Display()
	// Output: 4
}

func ExampleAVLInsertList() {
	a := NewAVL()
	a.InsertList(u.NewInteger(5), u.NewInteger(-4), u.NewInteger(32))
	a.Display()
	// Output: -4
	// 5
	// 32
}

func ExampleAVLCollision() {
	a := NewAVL()
	a.InsertList(u.NewInteger(5), u.NewInteger(-4), u.NewInteger(5))
	a.Display()
	// Output: -4
	// 5
}

func ExampleLeftRight() {
	a := NewAVL()
	a.InsertList(u.NewInteger(5), u.NewInteger(-5), u.NewInteger(0))
	a.Display()
	// Output: -5
	// 0
	// 5
}

func ExampleRightLeft() {
	a := NewAVL()
	a.InsertList(u.NewInteger(-5), u.NewInteger(5), u.NewInteger(0))
	a.Display()
	// Output: -5
	// 0
	// 5
}
func TestLeftLeft(t *testing.T) {
	a := NewAVL()
	a.InsertList(u.NewInteger(5), u.NewInteger(0), u.NewInteger(-5))
	if a.head.value.Compare(u.NewInteger(0)) != 0 {
		t.Fatal("Head should be 0")
	} else if a.head.left.value.Compare(u.NewInteger(-5)) != 0 {
		t.Fatal("Left should be -5")
	} else if a.head.right.value.Compare(u.NewInteger(5)) != 0 {
		t.Fatal("Right should be 5")
	}
}

func TestRightRight(t *testing.T) {
	a := NewAVL()
	a.InsertList(u.NewInteger(-5), u.NewInteger(0), u.NewInteger(5))
	if a.head.value.Compare(u.NewInteger(0)) != 0 {
		t.Fatal("Head should be 0")
	} else if a.head.left.value.Compare(u.NewInteger(-5)) != 0 {
		t.Fatal("Left should be -5")
	} else if a.head.right.value.Compare(u.NewInteger(5)) != 0 {
		t.Fatal("Right should be 5")
	}
}

func TestRightLeft(t *testing.T) {
	a := NewAVL()
	a.InsertList(u.NewInteger(-5), u.NewInteger(5), u.NewInteger(0))
	if a.head.value.Compare(u.NewInteger(0)) != 0 {
		t.Fatal("Head should be 0")
	} else if a.head.left.value.Compare(u.NewInteger(-5)) != 0 {
		t.Fatal("Left should be -5")
	} else if a.head.right.value.Compare(u.NewInteger(5)) != 0 {
		t.Fatal("Right should be 5")
	}
}

func TestLeftRight(t *testing.T) {
	a := NewAVL()
	a.InsertList(u.NewInteger(5), u.NewInteger(-5), u.NewInteger(0))
	if a.head.value.Compare(u.NewInteger(0)) != 0 {
		t.Fatal("Head should be 0")
	} else if a.head.left.value.Compare(u.NewInteger(-5)) != 0 {
		t.Fatal("Left should be -5")
	} else if a.head.right.value.Compare(u.NewInteger(5)) != 0 {
		t.Fatal("Right should be 5")
	}
}

func ExampleAVLLeftLeft() {
	a := NewAVL()
	a.InsertList(u.NewInteger(10), u.NewInteger(-4), u.NewInteger(-5))
	a.Display()
	// Output: -5
	// -4
	// 10
}
