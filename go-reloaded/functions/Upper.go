package goReloaded

func Upper(word string) string {
	res := ""
	for _, ele := range word {
		if ele >= 'A' && ele <= 'Z' {
			res += string(ele)
		} else if ele >= 'a' && ele <= 'z' {
			res += string(ele - 32)
		} else {
			res += string(ele)
		}
	}
	return res
}
