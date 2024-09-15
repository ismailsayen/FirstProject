package goReloaded

func isPunctuation(s rune) bool {
	if s == ',' || s == '?' || s == '!' || s == '.' || s == ':' || s == ';' {
		return true
	}
	return false
}
