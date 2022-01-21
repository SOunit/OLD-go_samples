package main

import "fmt"

func main() {
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	var c []int
	fmt.Println(c)

	c = append(c, 10)
	c = append(c, 20)
	c = append(c, 30)
	fmt.Println(c)

	d := []int{1, 2, 3}
	fmt.Println(d)
}
