package piscine

func Sqrt(nb int) int {
	if nb <= 3 {
		return 0
	}
	if nb == 4 {
		return 2
	}
	for i := 1; i <= nb/i; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
