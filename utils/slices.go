package utils

import "fmt"

func SlicesDriver() {
	// Make a slice of strings of length 3
	s := make([]string, 3)
	fmt.Println("set: ", s)

	s[0], s[1], s[2] = "a", "b", "c"

	fmt.Println("get: ", s[0])
	fmt.Println("len: ", len(s))

	fmt.Println("set: ", s)

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("set: ", s)

	// make a new slice of the same length as s
	c := make([]string, len(s))

	// copy the contents of s into c
	copy(c, s)
	fmt.Println("copy: ", c)

	// "slice", into a slice
	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2: ", l)

	// inline initialization of a slice
	t := []string{"a", "b", "c"}
	fmt.Println("init: ", t)
	fmt.Printf("%T\n", t)

	// create a 2D slice
	twoD := make([][]int, 3)

	// The length of the inner slices can vary, unlike a typical array
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
