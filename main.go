package main

import "fmt"

func main() {
	fmt.Println(greet("Raz"))
	fmt.Println(describe("Raz", 16, 1.84, false))
	fmt.Println(analyze([]int{1, 2, 3, 4}))
	arrayteste := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	arraya := []int{1, 2, 3, 4}
	arrayb := []int{2, 4, 6}
	fmt.Println(unique(arrayteste))
	fmt.Println(intersect(arraya, arrayb))
}
