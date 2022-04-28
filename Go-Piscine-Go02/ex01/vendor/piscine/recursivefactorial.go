package piscine

func RecursiveFactorial(nb int) int {
	const intSize = 32 << (^uint(0) >> 63)
	if nb < 0 {
		return 0
	}
	if (intSize == 64 && nb >= 21) || (intSize == 32 && nb >= 13) {
		return 0
	}
	if nb == 0 {
		return 1
	}
	return nb * RecursiveFactorial(nb-1)
}
