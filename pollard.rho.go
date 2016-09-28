package primer

func gcd(a, b int) int {
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}

// Implementation of Pollard's rho algorithm - An integer factorization algorithm
// Based on Floyd's cycle finding algorithm
func PollardRho(n int) int {
	xStatic := 2
	cycleSize := 2
	x := 2
	factor := 1

	for factor == 1 {
		for c := 1; (c <= cycleSize) && factor <= 1; c++ {
			x = (x*x + 1) % n
			factor = gcd(x-xStatic, n)
		}
		cycleSize *= 2
		xStatic = x
	}
	return factor
}
