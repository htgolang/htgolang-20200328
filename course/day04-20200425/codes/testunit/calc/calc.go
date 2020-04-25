package calc

func Add(n, m int) int {
	return n + m
}

func Multi(n, m int) int {
	return n * m
}

func Flag(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	} else {
		return 0
	}
}
