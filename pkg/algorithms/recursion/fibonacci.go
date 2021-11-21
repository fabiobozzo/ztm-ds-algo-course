package recursion

// O(n)
func fibonacciIterative(n int) []int {
	if n <= 1 {
		return []int{0}
	}

	series := []int{0, 1}
	first := 0
	second := 1

	for i := 2; i < n; i++ {
		next := first + second
		first = second
		second = next
		series = append(series, next)
	}

	return series
}

// O(2^n) - more readable but not performant
func fibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
