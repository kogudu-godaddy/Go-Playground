package utils

import "fmt"

func ArrayDriver() {
	// create array of ints of size 5
	var a [5]int
	// initialize the last variable of the array to 100
	a[4] = 100

	fmt.Print(a)
	fmt.Printf(" size: %v\n", len(a))

	// inline initialization of an array
	b := [5]int{1, 2, 3, 4, 5}

	fmt.Print(b)
	fmt.Printf(" size: %v\n", len(b))

	// create a 2D array 2 x 3
	var c [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println(c)
}
