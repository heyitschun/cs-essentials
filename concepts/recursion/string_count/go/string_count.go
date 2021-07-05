package stringcount

func StringCount(a []string, c int) int {
	sub := a[0]
	next := a[1:]
	c += len(sub)
	if len(next) != 0 {
		return StringCount(next, c)
	}
	return c
}
