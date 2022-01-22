package main

import "fmt"

func main() {
	i := 7
	fmt.Println(i)
	fmt.Println(&i)

	inc(i)
	fmt.Println(i)

	incPointer(&i)
	fmt.Println(i)

}

func inc(x int) {
	x++
}

func incPointer(x *int) {
	*x++
}
