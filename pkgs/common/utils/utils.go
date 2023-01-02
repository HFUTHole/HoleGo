package utils

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}
