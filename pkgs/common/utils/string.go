package utils

import "strings"

func Scape(str string) string {
	replace := strings.Replace(str, "<", "&lt;", -1)
	replace = strings.Replace(replace, ">", "&gt;", -1)
	//replace = strings.Replace(replace, " ", "&nbsp;", -1)
	return replace
}

func ScapeSlice(str []string) []string {
	res := make([]string, len(str))
	for i, s := range str {
		replace := strings.Replace(s, "<", "&lt;", -1)
		replace = strings.Replace(replace, ">", "&gt;", -1)
		replace = strings.Replace(replace, " ", "&nbsp;", -1)
		res[i] = replace
	}
	return res
}

func SliceElementMaxLength(slice []string, max int) bool {
	for _, s := range slice {
		if len(s) >= max {
			return false
		}
	}
	return true
}
