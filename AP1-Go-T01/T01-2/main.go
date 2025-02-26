package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

type Pair struct {
	Key   string
	Value int
}

func main() {
	uniqueWords := make(map[string]int)
	var words []string
	var number uint64

	pairs := make([]Pair, 0, len(uniqueWords))
	
	if readString(&words, &number) < 1 || len(words) < 1 {
		fmt.Printf("\n")
	} else {
		uniqWordsCount(uniqueWords, words)
		createPairSlice(&pairs, uniqueWords)
		sortWords(&pairs)
		selectWordCount(&pairs, number)
	}

}

func readString(words *[]string, number *uint64) uint64 {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	*words = strings.Split(scanner.Text(), " ")

	for {
		scanner.Scan()
		k := scanner.Text()

		var err error
		*number, err = strconv.ParseUint(k, 10, 8)

		if err != nil {
			fmt.Println("Введи положительное число")
		} else {
			break
		}
	}

	return *number
}

func uniqWordsCount(uniqueWords map[string]int, words []string) {
	for _, el := range words {
		_, ok := uniqueWords[el]
		if ok {
			uniqueWords[el]++
		} else {
			uniqueWords[el] = 1
		}
	}
}

func createPairSlice(pairs *[]Pair, uniqueWords map[string]int) {
	for key, value := range uniqueWords {
		*pairs = append(*pairs, Pair{key, value})
	}
}

func sortWords(pairs *[]Pair) {
	sort.Slice(*pairs, func(i, j int) bool {
		if (*pairs)[i].Value != (*pairs)[j].Value {
			return (*pairs)[i].Value > (*pairs)[j].Value
		}
		return (*pairs)[i].Key < (*pairs)[j].Key
	})
}

func selectWordCount(pairs *[]Pair, number uint64) {
	if len(*pairs) < int(number) {
		number = uint64(len(*pairs))
		// printWords(pairs, len(*pairs))
	}
	printWords(pairs, int(number))
}

func printWords(pairs *[]Pair, number int) {
	for i := 0; i < number; i++ {
		fmt.Print((*pairs)[i].Key)
		if i != number - 1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}