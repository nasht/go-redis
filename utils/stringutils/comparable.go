package stringutils

func Pop(a string) (string, string) {
	if len(a) == 0 {
		return "", ""
	}
	return string(a[len(a)-1]), string(a[:len(a)-1])
}

func Shift(a string) (string, string) {
	if len(a) == 0 {
		return "", ""
	}
	return string(a[0]), string(a[1:])
}