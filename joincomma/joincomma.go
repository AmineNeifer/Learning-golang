package joincomma

import "strings"

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0 {
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		if len(phrases) != 2 {
			result += ","
		}
		result += " and "
		result += phrases[len(phrases)-1]
		return result
	}

}
