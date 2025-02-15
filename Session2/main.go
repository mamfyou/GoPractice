// package main

// var a int = 1

// func addNumber(a *int) (res int) {
// 	*a += 1
// 	return *a
// }

// func main() {
// 	addNumber(&a)
// 	print(a)
// }

package main

import (
    "fmt"
    "time"
)

func main() {
    currentTime := time.Now()
	fmt.Println(currentTime)
	for i := 0; i < 1_000_000_000_000; i++ {

	}
	new_time := time.Now()
	duration := new_time.Sub(currentTime)
    fmt.Printf("Mili Seconds: %d\n", int(duration.Milliseconds()))
}