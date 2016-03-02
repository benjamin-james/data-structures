package text

import (
	u "benJames/util"
	"bufio"
	"os"
	"strings"
)

func Clean(str string) string {
	//lowercase, remove punctuation, numbers
	str = strings.ToLower(str)
	return str
}

func ReadFile(filename string, unigrams, bigrams u.DataStructure) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	prev := ""
	for scanner.Scan() {
		str := scanner.Text()
		str = Clean(str)
		unigrams.Insert(NewUni(str, 1))
		if prev != "" {
			bigrams.Insert(NewBi(prev, str, 1))
		}
		prev = str
	}
	return scanner.Err()
}
