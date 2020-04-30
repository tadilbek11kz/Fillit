package utils

func Min(arr []int) int {
	sm := 1<<63 - 1
	for _, nbr := range arr {
		if nbr < sm {
			sm = nbr
		}
	}
	return sm
}

func Max(arr []int) int {
	sm := -1
	for _, nbr := range arr {
		if sm < nbr {
			sm = nbr
		}
	}
	return sm
}

func ValidateTouchingSides(sides int) bool {
	if sides == 6 || sides == 8 {
		return true
	}
	panic("ERROR")
}

func CountTouchingSides(shape [][]byte, i, j, length int) int {
	count := 0
	if 0 <= j-1 && shape[j-1][i] == '#' {
		count++
	}
	if j+1 < length && shape[j+1][i] == '#' {
		count++
	}
	if 0 <= i-1 && shape[j][i-1] == '#' {
		count++
	}
	if i+1 < length && shape[j][i+1] == '#' {
		count++
	}
	return count
}
