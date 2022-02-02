package utils

import "fmt"

func FlowControlDriver() {
	const num = 10
	fizzBuzz(num)
	if num2 := num; num2%2 == 0 {
		fmt.Println("Second variable is EVEN")
	} else {
		fmt.Println("Second variable is ODD")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean")
		case int:
			fmt.Println("Integer")
		default:
			fmt.Printf("%T\n", t)
		}
	}

	whatAmI(3)
	whatAmI("Kelechi")
	whatAmI(false)
}

func fizzBuzz(n int) {
	for i := n; i >= 0; i-- {
		if i%2 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println("buzz")
		}
	}
}
