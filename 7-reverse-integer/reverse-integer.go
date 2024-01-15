func reverse(x int) int {
	var rev int
	var isNeg bool

	if x < 0 {
		isNeg = true
		x = x * (-1)
	}

	for x > 0 {
		rev = rev*10 + x%10
		x = x / 10
	}

	if isNeg {
		rev = rev * (-1)
	}

     if rev < Pow(2, 31)*(-1) || rev > Pow(2, 31) - 1 {
        return 0
    }

	return rev
}

func Pow(base, exponent int) int {
	result := 1

	for i := 0; i < exponent; i++ {
		result *= base
	}

	return result
}