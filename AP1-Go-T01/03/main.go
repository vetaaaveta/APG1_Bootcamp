package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var num1, num2 []int
	var value = make(map[int]int)

	if readStr(&num1) == nil && readStr(&num2) == nil {
		searchMatch(&num1, &num2, value)
	} else {
		fmt.Println("Invalid input")
	}
}

func printMatch(num1 *[]int, value map[int]int) {
	if len(value) > 0 {
		for i, el := range *num1 {
			_, ok := value[el]
			if ok {
				fmt.Print(el)
				if i < len(*num1) - 1 {
					fmt.Print(" ")
				} else {
					fmt.Print("\n")
				}
			}
		}
	} else {
		fmt.Println("Empty intersection")
	}
}

func searchMatch(num1 *[]int, num2 *[]int, value map[int]int) {
	for i := 0; i < len(*num1); i++ {
		for j := 0; j < len(*num2); j++ {
			if (*num1)[i] == (*num2)[j] {
				value[(*num1)[i]]++
			}
		}
	}
	printMatch(num1, value)
}

func readStr(numbers *[]int) error {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var err error
	for _, el := range scanner.Text() {
		num, err := strconv.Atoi(string(el))
		if err == nil {
			*numbers = append(*numbers, num)
		}
	}
	return err
}