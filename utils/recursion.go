package utils

func Factorial(n int) int {
	if n < 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonnaci(n int) int {
	var fib func(n int) int

	memo := make(map[int]int)

	fib = func(n int) int {
		if val, ok := memo[n]; ok {
			return val
		}

		if n < 2 {
			return n
		}
		memo[n] = fib(n-1) + fib(n-2)
		return fib(n-1) + fib(n-2)
	}

	return fib(n)
}
