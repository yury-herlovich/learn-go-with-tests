package iteration

func Repeat(symbol string, c int) string {
	res := ""

	for i := 0; i < c; i++ {
		res += symbol
	}

	return res
}
