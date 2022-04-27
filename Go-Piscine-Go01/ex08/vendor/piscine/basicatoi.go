package piscine

func BasicAtoi(s string) int {
	res := 0
	for _, v := range s {
		if v < '0' || v > '9' {
			break
		}
		res = res*10 + int(byte(v)-'0')
	}
	return res
}
