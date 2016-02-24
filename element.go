package main

type element interface {
	Compare(element) int
	Update()
	String() string
}
