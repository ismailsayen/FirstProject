package goReloaded

func Split(str string) []string {
	slice := []string{}
	word := ""
	cha := ""
	ord := false
	d := ""
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "(" {
			ord = true
			continue
		} else if string(str[i]) == ")" {
			ord = false
			continue
		}
		if ord {
			d += string(str[i])
		} else {
			if d != "" {
				slice = append(slice, "("+d+")")
				d = ""
			}
		}
		if !ord {
			if string(str[i]) != " " && string(str[i]) != "\t" && !isPunctuation(string(str[i])) {
				if cha != "" {
					slice = append(slice, cha)
					cha = ""
				}
				word += string(str[i])
			} else {
				if word != "" {
					slice = append(slice, word)
					word = ""
				}
				if isPunctuation(string(str[i])) {
					cha += string(str[i])
				} else {
					if cha != "" {
						slice = append(slice, cha)
						cha = ""
					}
				}
			}
		}
	}
	if word != "" {
		slice = append(slice, word)
	}
	if cha != "" {
		slice = append(slice, cha)
	}
	return slice
}

