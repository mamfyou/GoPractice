package main

import "fmt"

func drawPyramid(n int) {
	num := 1
	max_space := 2 * n - 1
	
	for i := n; i > 0; i-- {
		for space_count := 1; space_count < max_space; space_count++ {
			fmt.Print(" ")
		}
		max_space -= 2

		for ascending_number := 1; ascending_number <= num; ascending_number++ {
			fmt.Print(ascending_number)
			fmt.Print(" ")
		}

		for descending_number := num - 1; descending_number >= 1; descending_number-- {
			fmt.Print(descending_number)
			fmt.Print(" ")
		}

		fmt.Print("\n")
		num++
	}
}

func main() {
    var n int
    fmt.Scanf("%d", &n)
    
	drawPyramid(n)
}