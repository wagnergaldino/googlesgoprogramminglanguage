package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(x ...int) {
	for i, v := range x {
		fmt.Println("indice", i, "valor", v)
	}
	fmt.Println("----------------------------")
}
