package main

import (
	"bufio"
	"os"
)

type datastructure interface {
	Display()
	Insert(element)
	Find(element) element
}

func ReadFile(filename string, unigrams, bigrams datastructure) {
	file, _ := os.Open(filename)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	prev := ""
	for scanner.Scan() {
		str := scanner.Text()
		unigrams.Insert(NewUni(str, 1))
		if prev != "" {
			bigrams.Insert(NewBi(prev, str, 1))
		}
		prev = str
	}
}
