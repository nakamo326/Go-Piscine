package piscine

func IterativeFactorial(nb int) int {
	var intmax int = int(^uint(0) >> 1)
	res := 1
	if nb < 0 {
		return 0
	}
	for i := 1; i <= nb; i++ {
		if res >= intmax/i {
			return 0
		}
		res = res * i
	}
	return res
}
