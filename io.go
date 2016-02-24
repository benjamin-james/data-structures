package main

import (
	"bufio"
	"fmt"
	"os"
)

type datastructure interface {
	Display()
	Insert(element)
	Find(element) element
}

func ReadFile(filename string, unigrams, bigrams datastructure) error {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return err
	}
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
	return nil
}
