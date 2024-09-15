package goReloaded

func Split(str string) []string {
	slice := []string{}
	word := ""
	for _, r := range str {
		if r != ' ' && r != '\t' {
			word += string(r) 
		} else {
			if word != "" {
				slice = append(slice, word)
				word = ""
			}
		}
	}
	if word != "" {
		slice = append(slice, word)
	}

	return slice
}
