package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Array 1: ")
	scanner.Scan()
	input1 := scanner.Text()

	fmt.Println("Array 2 ")
	scanner.Scan()
	input2 := scanner.Text()

	input1 = strings.ReplaceAll(input1, " ", "")
	input2 = strings.ReplaceAll(input2, " ", "")

	if len(input1) != len(input2) {
		fmt.Println("Array tidak sama panjang")
	}

	var diffs []int

	for i := 0; i < len(input1); i++ {
		if input1[i] != input2[i] {
			diffs = append(diffs, i)
			fmt.Printf("Indeks ke: %v berbeda\n", i)
		}
	}

	if len(diffs) == 0 {
		fmt.Println("Array Sama")
	}

}
