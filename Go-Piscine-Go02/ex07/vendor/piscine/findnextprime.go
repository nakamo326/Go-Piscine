package piscine

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/i; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	i := nb
	for ; ; i++ {
		if isPrime(i) {
			break
		}
	}
	return i
}
