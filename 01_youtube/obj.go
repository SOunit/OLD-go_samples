package main

import "fmt"

func main() {
	vertices := make(map[string]int)
	vertices["test1"] = 1
	vertices["test2"] = 2
	vertices["test3"] = 3

	fmt.Println(vertices)
	fmt.Println(vertices["test1"])

	delete(vertices, "test1")
	fmt.Println(vertices)

}
