package gopkg

func Sum(x ...int) (res int) {
	for _, v := range x {
		res += v
	}
	return
}
