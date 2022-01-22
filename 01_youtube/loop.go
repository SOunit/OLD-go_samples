package main

import "fmt"

func main() {
	// for loop
	fmt.Println("------------------ i")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("------------------ j")

	// while loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// array loop
	fmt.Println("--------------- arr")
	arr := []int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Print("index", index)
		fmt.Println("value", value)
	}

	// obj loop
	fmt.Println("--------------- obj")
	obj := make(map[string]string)
	obj["j"] = "Jack"
	obj["r"] = "Rebecca"
	fmt.Println(obj)

	for key, value := range obj {
		fmt.Println("key", key)
		fmt.Println("value", value)
	}
}
