package sum

func Add(start int, max int) (sum int) {
	var i int
	for i = start; i < max; i++ {
		sum += i
	}
	return
}
