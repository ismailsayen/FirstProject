package goReloaded

func LowerCase(word string) string {
	res := ""
	for _, ele := range word {
		if ele >= 'a' && ele <= 'z' {
			res += string(ele)
		} else if ele >= 'A' && ele <= 'Z' {
			res += string(ele + 32)
		} else {
			res += string(ele)
		}
	}
	return res
}
