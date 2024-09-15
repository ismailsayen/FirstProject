package goReloaded

func SplitWithPunc(slice []string) []string {
	new_slice := []string{}
	res := ""
	a := ""
	p := ""
	for i, ele := range slice {
		if i != len(slice)-1 {
			res += ele + " "
		} else {
			res += ele
		}
	}
	for _, e := range res {
		if isPunctuation(e) {
			p += string(e)
		}
		if e != ' ' && e != '\t' && !isPunctuation(e) {
			if p != "" {
				new_slice = append(new_slice, p)
				p = ""
			}
			a += string(e)
		} else {
			if a != "" {
				new_slice = append(new_slice, a)

				a = ""
			}
		}
	}
	if a != "" {
		new_slice = append(new_slice, a)
	}
	if p != "" {
		new_slice = append(new_slice, p)
	}
	return new_slice
}