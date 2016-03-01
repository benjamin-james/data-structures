package util

type Element interface {
	Compare(Element) int
	Update()
	String() string
}

type DataStructure interface {
	Display()
	Insert(Element)
	Find(Element) Element
	Iterator() <-chan Element
}
