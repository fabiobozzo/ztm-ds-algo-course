package recursion

func RecursiveFactorial(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return n * RecursiveFactorial(n-1)
}

func IterativeFactorial(n int64) int64 {
	var result int64
	result = 1

	for i := n; i >= 1; i-- {
		result = i * result
	}

	return result
}
