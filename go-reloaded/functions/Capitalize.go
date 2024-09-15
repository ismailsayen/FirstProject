package goReloaded

func prim(a rune) bool {
	return a >= 'A' && a <= 'Z' || a >= 'a' && a <= 'z' || a >= '0' && a <= '9'
}

func Capitalize(s string) string {
	ar := []rune(s)
	letra := true
	for i := 0; i < len(ar); i++ {
		if prim(ar[i]) && letra {
			if ar[i] >= 'a' && ar[i] <= 'z' {
				ar[i] = ar[i] - ('a' - 'A')
			}
			letra = false
		} else if ar[i] >= 'A' && ar[i] <= 'Z' {
			ar[i] = ar[i] + ('a' - 'A')
		} else if !prim(ar[i]) {
			letra = true
		}
	}
	return string(ar)
}
