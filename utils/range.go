package utils

import "fmt"

/*
 range can be used to iterate over elements in a variety of data
 structrues.
*/
func RangeDriver() {

	// using range to iterate over a slice
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum of 1 + 2 + 3 = ", sum)

	// using range to iterate over a map
	m := map[string]string{"a": "apple", "b": "ball"}

	for k, v := range m {
		fmt.Println(k, " -> ", v)
	}

	for k := range m {
		fmt.Println("key: ", k)
	}

	// using range to iterate over a string
	for i, c := range "kelechi" {
		fmt.Println(i, ": ", string(c))
	}
}
