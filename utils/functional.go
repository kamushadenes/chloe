package utils

func SubtractInt(ints ...int) int {
	var result int
	for _, i := range ints {
		result -= i
	}
	return result
}

func SubtractIntWithMinimum(min int, ints ...int) int {
	result := SubtractInt(ints...)

	if result < min {
		return min
	}

	return result
}
