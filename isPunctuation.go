package goReloaded

func isPunctuation(s string) bool {
	if s == "," || s == "?" || s == "!" || s == "." || s == ":" || s == ";" {
		return true
	}
	return false
}
