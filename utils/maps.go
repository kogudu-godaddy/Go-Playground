package utils

import "fmt"

func MapDriver() {

	// create a map [string] -> int
	m := make(map[string]int)

	m["kelechi"] = 1
	m["emeka"] = 2
	m["chimere"] = 3

	fmt.Println(m, "len: ", len(m))
	fmt.Println(m["emeka"])

	// delete map key
	delete(m, "kelechi")
	fmt.Println(m, "len: ", len(m))

	// initalize map inline
	n := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println(n, "len: ", len(n))
}
