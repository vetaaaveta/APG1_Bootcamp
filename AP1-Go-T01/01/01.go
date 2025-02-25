package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	text := []string{"Введите первое число", "Введите оператор", "Введите второе число"}
	var nums []float64

	operator := read(text, &nums)
	value, err := calc(operator, nums)
	
	if err == nil {
		clearConsole()
		fmt.Printf("Результат: %.3f\n", value)
	} else {
		clearConsole()
		fmt.Println(err)
	}

}

func read(text []string, nums *[]float64) string {
	scanner := bufio.NewScanner(os.Stdin)
	var operator string
	for i := 0; i < 3; i++ {
		clearConsole()
		fmt.Println(text[i])
		scanner.Scan()
		if i%2 == 0 {
			num, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				repeatInput(&i)
			} else {
				*nums = append(*nums, num)
			}
		} else {
			operator = scanner.Text()
			if !validOperator(operator) {
				repeatInput(&i)
			}
		}
	}
	return operator
}

func clearConsole() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func repeatInput(i *int) {
	delay := 500 * time.Millisecond
	time.Sleep(delay)
	clearConsole()
	fmt.Println("Invalid input")
	time.Sleep(delay*2)
	clearConsole()
	*i--

}

func validOperator(operator string) bool {
	flag := false
	switch operator {
	case "+":
		flag = true
	case "-":
		flag = true
	case "*":
		flag = true
	case "/":
		flag = true
	}
	return flag
}

func calc(operator string, nums []float64) (float64, error) {
	var value float64
	var err error

	switch operator {
	case "+":
		value = nums[0] + nums[1]
	case "-":
		value = nums[0] - nums[1]
	case "*":
		value = nums[0] * nums[1]
	case "/":
		if nums[1] != 0 {
			value = nums[0] / nums[1]
		} else {
			err = errors.New("НА НОЛЬ ДЕЛИТЬ НЕЛЬЗЯ!!!!!!!!!!!!!!!!!!")
		}
	}

	return value, err
}
