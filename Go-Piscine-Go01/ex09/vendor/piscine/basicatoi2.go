package piscine

func BasicAtoi2(s string) int {
	res := 0
	for _, v := range s {
		if v < '0' || v > '9' {
			return 0
		}
		res = res*10 + int(byte(v)-'0')
	}
	return res
}
