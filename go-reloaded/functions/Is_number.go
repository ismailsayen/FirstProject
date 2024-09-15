package goReloaded

func Is_number(str string) bool {
	for _, e := range str {
		if e < '0' && e > '9' {
			return false
		}
	}
	return true
}
