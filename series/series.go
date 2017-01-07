package slice

func All(n int, txt string) []string {
	subs := []string{}
	for i := 0; i <= len(txt)-n; i++ {
		subs = append(subs, txt[i:i+n])
	}
	return subs
}

func UnsafeFirst(n int, txt string) string {
	return txt[:n]
}

func First(n int, txt string) (string, bool) {
	if n > len(txt) {
		return "", false
	}

	return txt[:n], true
}
