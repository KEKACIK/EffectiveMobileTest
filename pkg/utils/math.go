package utils

func RoundUp(a, b int) int {
	return (a + b - 1) / b
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
