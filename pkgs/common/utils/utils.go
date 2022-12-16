package utils

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IntToBoll(i int) bool {
	if i > 0 {
		return true
	}
	return false
}
