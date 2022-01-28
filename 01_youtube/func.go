package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sum(5, 5))
	fmt.Println(sum(5, 10))

	val := 0

	result, err := calc(val)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sum(num1 int, num2 int) int {
	return (num1 + num2)
}

func calc(num int) (int, error) {
	if num <= 0 {
		return num, errors.New("Undefined for negative numbers")
	}

	return num * 2, nil
}
