package text

import (
	u "benJames/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Clean(str string) string {
	//lowercase, remove punctuation, numbers
	str = strings.ToLower(str)
	return str
}

func WriteString(filename string, s string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, s)
	return err
}
func ReadFile(filename string, unigrams, bigrams u.DataStructure) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	prev := ""
	count := 0
	for scanner.Scan() {
		str := scanner.Text()
		str = Clean(str)
		unigrams.Insert(NewUni(str, 1))
		if prev != "" {
			bigrams.Insert(NewBi(prev, str, 1))
		}
		prev = str
		count++
	}
	return count, scanner.Err()
}

func DumpToFile(filename string, ds u.DataStructure) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	ds.Display(file)
	return nil
}
