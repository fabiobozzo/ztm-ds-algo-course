package recursion

func ReverseString(s string) string {
	chars := []rune(s)

	swapHeadAndTail(chars, 0)

	return string(chars)
}

func swapHeadAndTail(chars []rune, offset int) {
	rightOffset := len(chars) - 1 - offset

	if offset >= len(chars)/2 {
		return
	}

	chars[offset], chars[rightOffset] = chars[rightOffset], chars[offset]

	swapHeadAndTail(chars, offset+1)
}
