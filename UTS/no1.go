package main

import "fmt"

func sums(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Aku", "suka", "makan", "Bubur Ayam")

	fmt.Println(sums(1, 2, 3, 4, 5, 6))

	nums := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sums(nums...))

}
