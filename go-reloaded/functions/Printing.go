package goReloaded

func Printing(slice []string) string {
	res := ""
	loop := func(s string) bool {
		for _, e := range s {
			return isPunctuation(e)
		}
		return false
	}

	for i, ele := range slice {
		if loop(ele) {
			continue
		}
		if i < len(slice)-1 {
			if loop(slice[i+1]) {
				res += ele + slice[i+1] + " "
			} else {
				res += ele + " "
			}
		} else {
			res += ele
		}
	}
	if len(res) == 0 {
		return "\n"
	}
	if res[len(res)-1] == ' ' {
		res = res[0 : len(res)-1]
	}
	return res +"\n"
}
