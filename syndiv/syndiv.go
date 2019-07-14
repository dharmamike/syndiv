package syndiv

func Syndiv(coeffs []int, c int) []int {
	n := len(coeffs)

	if n == 0 {
		return []int{}
	}

	quotient := make([]int, n)
	quotient[0] = coeffs[0]

	for i := 0; i < n-1; i++ {
		quotient[i+1] = coeffs[i+1] + c*quotient[i]
	}

	return quotient
}
