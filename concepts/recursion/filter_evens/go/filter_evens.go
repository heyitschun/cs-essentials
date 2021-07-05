package filter

func FilterEvens(a []int) []int {
	var ret []int
	for _, n := range a {
		if n%2 == 0 {
			ret = append(ret, n)
		}
	}

	return ret
}
