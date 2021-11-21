package arrays

func reverse(s string) string {
	chars := []rune(s)

	i := 0
	j := len(chars) - 1

	for i < j {
		chars[i], chars[j] = chars[j], chars[i]

		i++
		j--
	}

	return string(chars)
}
